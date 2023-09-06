package logic

import (
	"context"
	"douyin/rpc/interactive/internal/JWT"
	"douyin/rpc/interactive/internal/svc"
	"douyin/rpc/interactive/pb"
	"douyin/rpc/interactive/utils"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinFavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinFavoriteListLogic {
	return &DouyinFavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinFavoriteListLogic) DouyinFavoriteList(in *pb.DouyinFavoriteListRequest) (*pb.DouyinFavoriteListResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("进入 rpc DouyinFavoriteList")
	if err := JWT.JWTAuthToken(l.svcCtx, l.ctx, in.Token); err != nil {
		fmt.Println("token鉴权出错了")
		return &pb.DouyinFavoriteListResponse{
			StatusCode: 1,
			StatusMsg:  "token鉴权出错了",
		}, err
	}
	// 根据user_id从favorite中查找video
	favoriteList, err := l.svcCtx.FavoriteModel.FindByUser(l.ctx, in.UserId)
	if err != nil {
		return &pb.DouyinFavoriteListResponse{
			StatusCode: 1,
			StatusMsg:  "从favorite中根据用户id查找视频失败",
		}, err
	}
	var VideoList []*pb.Video
	for _, favorite := range favoriteList {
		// 从favorite中获取video_id
		mvideo, err := l.svcCtx.VideoModel.FindOne(l.ctx, favorite.VideoId.Int64)
		if err != nil {
			return &pb.DouyinFavoriteListResponse{
				StatusCode: 1,
				StatusMsg:  "从Video中根据视频id查找视频失败",
			}, err
		}
		// model.video中存放author的id
		mauthor, err := l.svcCtx.UserModel.FindOne(l.ctx, mvideo.Author)
		if err != nil {
			return &pb.DouyinFavoriteListResponse{
				StatusCode: 1,
				StatusMsg:  "根据video中的author id查找author失败",
			}, err
		}
		author, _ := utils.UserModelPb(mauthor)
		video, _ := utils.VideoModelPb(mvideo, author)
		video.IsFavorite = true
		VideoList = append(VideoList, video)
	}
	fmt.Println(VideoList)
	return &pb.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "成功",
		VideoList:  VideoList,
	}, nil
}
