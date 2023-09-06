package feed

import (
	"context"
	"douyin/api/core/internal/svc"
	"douyin/api/core/internal/types"
	"douyin/rpc/core/pb"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinFeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinFeedLogic {
	return &DouyinFeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinFeedLogic) DouyinFeed(req *types.DouyinFeedReq) (resp *types.DouyinFeedResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println(req.Latest_time)

	rpcReq := pb.DouyinFeedRequest{
		LatestTime: req.Latest_time,
		Token:      req.Token,
	}
	rpcResp, err := l.svcCtx.CoreRpcClient.DouyinFeed(l.ctx, &rpcReq)
	if err != nil {
		return &types.DouyinFeedResp{
			Status_code: 1,
		}, err
	}

	var videoList []types.Video
	for _, video := range rpcResp.VideoList {
		thisAuthor := types.User{
			Id:             video.Author.Id,
			Name:           video.Author.Name,
			Follow_count:   video.Author.FollowCount,
			Follower_count: video.Author.FollowerCount,
			Is_follow:      video.Author.IsFollow,
		}
		thisVideo := types.Video{
			Id:             video.Id,
			Author:         thisAuthor,
			Play_url:       video.PlayUrl,
			Cover_url:      video.CoverUrl,
			Favorite_count: video.FavoriteCount,
			Comment_count:  video.CommentCount,
			Is_favorite:    video.IsFavorite,
			//Title:          video.Title,
		}
		fmt.Println(thisVideo)
		videoList = append(videoList, thisVideo)
	}
	return &types.DouyinFeedResp{
		Status_code: 0,
		Next_time:   rpcResp.NextTime,
		Video_list:  videoList,
	}, nil
}
