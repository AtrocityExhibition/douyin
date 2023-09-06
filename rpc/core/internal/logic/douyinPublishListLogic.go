package logic

import (
	"context"
	"douyin/rpc/core/internal/JWT"
	"douyin/rpc/core/internal/svc"
	"douyin/rpc/core/pb"
	"douyin/rpc/core/utils"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinPublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinPublishListLogic {
	return &DouyinPublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinPublishListLogic) DouyinPublishList(in *pb.DouyinPublishListRequest) (*pb.DouyinPublishListResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("进入 rpc DouyinPublishList")
	err := JWT.JWTAuthToken(l.svcCtx, l.ctx, in.Token)
	if err != nil {
		fmt.Println("用户鉴权错误")
		return &pb.DouyinPublishListResponse{
			StatusCode: 1,
		}, err
	}
	// 1.找到该用户
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		fmt.Println("根据user_id查询用户出错")
		return &pb.DouyinPublishListResponse{
			StatusCode: 1,
		}, err
	}
	// 2.根据用户id去video里找到该用户的videoList
	mVideoList, err := l.svcCtx.VideoModel.FindAllByUser(l.ctx, user.Id)
	if err != nil {
		fmt.Println("查找用户的video_list出错")
		return &pb.DouyinPublishListResponse{
			StatusCode: 1,
		}, err
	}
	var videoList []*pb.Video
	for _, mVideo := range mVideoList {
		author, _ := utils.UserModelPb(user)
		video, _ := utils.VideoModelPb(mVideo, author)
		videoList = append(videoList, video)
	}
	return &pb.DouyinPublishListResponse{
		StatusCode: 0,
		VideoList:  videoList,
	}, nil
}
