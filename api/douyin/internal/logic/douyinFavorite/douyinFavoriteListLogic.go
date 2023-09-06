package douyinFavorite

import (
	"context"
	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"douyin/rpc/interactive/pb"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinFavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinFavoriteListLogic {
	return &DouyinFavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinFavoriteListLogic) DouyinFavoriteList(req *types.DouyinFavoriteListRequest) (resp *types.DouyinFavoriteListResponse, err error) {
	// todo: add your logic here and delete this line

	rpcReq := pb.DouyinFavoriteListRequest{
		Token:  req.Token,
		UserId: req.User_id,
	}
	fmt.Println("调用 rpc DouyinFavoriteList")
	rpcResp, err := l.svcCtx.InteractiveRpcClient.DouyinFavoriteList(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("rpc 服务 DouyinFavoriteAction 出错了")
		return &types.DouyinFavoriteListResponse{
			Status_code: "1",
			Status_msg:  "rpc 服务 DouyinFavoriteAction 出错了",
		}, err
	}
	var VideoList []types.Video
	for _, pvideo := range rpcResp.VideoList {
		//thisAuthor := &types.User{
		//	Id:             pvideo.Author.Id,
		//	Name:           pvideo.Author.Name,
		//	Follow_count:   pvideo.Author.FollowCount,
		//	Follower_count: pvideo.Author.FollowerCount,
		//	Is_follow:      pvideo.Author.IsFollow,
		//}
		video := types.Video{
			Id: pvideo.Id,
			//Author:         thisAuthor,
			Play_url:       pvideo.PlayUrl,
			Cover_url:      pvideo.CoverUrl,
			Favorite_count: pvideo.FavoriteCount,
			Comment_count:  pvideo.CommentCount,
			Is_favorite:    pvideo.IsFavorite,
			Title:          pvideo.Title,
		}
		VideoList = append(VideoList, video)
		fmt.Println(video)
	}

	return &types.DouyinFavoriteListResponse{
		Status_code: "0",
		Video_list:  VideoList,
	}, nil
}
