package logic

import (
	"context"
	"douyin/rpc/core/internal/JWT"
	"douyin/rpc/core/internal/svc"
	"douyin/rpc/core/pb"
	"douyin/rpc/core/utils"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinFeedLogic {
	return &DouyinFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinFeedLogic) DouyinFeed(in *pb.DouyinFeedRequest) (*pb.DouyinFeedResponse, error) {
	// todo: add your logic here and delete this line
	videolist, err := InitVideoList(l.svcCtx, l.ctx)
	if err != nil {
		fmt.Println("查找全部视频出错")
	}
	// video 是否 点过赞需要去favorite中查看
	fmt.Println("token = ", in.Token)
	if in.Token != "" {
		claims, _ := JWT.JWTAuth(in.Token)
		username := claims["Username"].(string)
		password := claims["Password"].(string)
		u, err := l.svcCtx.UserModel.FindOneByToken(l.ctx, username, password)
		if err != nil {
			fmt.Println("查找用户错误")
		}
		// 根据u_id和video_id去favorite中查找，若存在则为点赞为true，否则为false
		// author中的is_follow也需要查表
		for _, v := range videolist {
			// 1.
			_, err := l.svcCtx.FavoriteModel.FindOneByUserVideo(l.ctx, u.Id, v.Id)
			if err != nil {
				v.IsFavorite = false
			} else {
				v.IsFavorite = true
			}
			// 2.
			_, err = l.svcCtx.Relationmodel.FindOneByUserToUser(l.ctx, u.Id, v.Author.Id)
			if err != nil {
				v.Author.IsFollow = false
			} else {
				v.Author.IsFollow = true
			}
		}
	}
	return &pb.DouyinFeedResponse{
		StatusCode: 0,
		VideoList:  videolist,
		NextTime:   time.Now().Unix(),
	}, nil
}

func InitVideoList(svcCtx *svc.ServiceContext, ctx context.Context) ([]*pb.Video, error) {
	// 如果传入了token，需要判断video的favorite，否则favorite全部为false
	// 从数据库中加载videolist
	var videolist = []*pb.Video{}

	result, err := svcCtx.VideoModel.FindAll(ctx)
	if err != nil {
		fmt.Println("查询所有video出错了")
		fmt.Println(err)
		return videolist, err
	}

	for _, video := range result {
		user, err := svcCtx.UserModel.FindOne(ctx, video.Author)
		if err != nil {
			fmt.Println("FindOneByToken 出错了")
			fmt.Println(err)
			return videolist, err
		}
		author, _ := utils.UserModelPb(user)
		pbVideo, _ := utils.VideoModelPb(video, author)
		videolist = append(videolist, pbVideo)
	}
	return videolist, nil
}
