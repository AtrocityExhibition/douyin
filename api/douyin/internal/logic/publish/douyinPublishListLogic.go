package publish

import (
	"context"
	"douyin/rpc/core/pb"
	"fmt"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinPublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinPublishListLogic {
	return &DouyinPublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinPublishListLogic) DouyinPublishList(req *types.DouyinPublishListReq) (resp *types.DouyinPublishListResp, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinPublishListRequest{
		Token:  req.Token,
		UserId: req.User_id,
	}

	rpcResp, err := l.svcCtx.CoreRpcClient.DouyinPublishList(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("调用 rpc DouyinPublishList 错误")
		return &types.DouyinPublishListResp{
			Status_code: rpcResp.StatusCode,
		}, err
	}

	var videoList []types.Video
	for _, video := range rpcResp.VideoList {
		thisVideo := types.Video{
			Id:             video.Id,
			Play_url:       video.PlayUrl,
			Cover_url:      video.CoverUrl,
			Favorite_count: video.FavoriteCount,
			Comment_count:  video.CommentCount,
			Is_favorite:    video.IsFavorite,
			Title:          video.Title,
		}
		fmt.Println(thisVideo)
		videoList = append(videoList, thisVideo)
	}
	return &types.DouyinPublishListResp{
		Status_code: rpcResp.StatusCode,
		Video_list:  videoList,
	}, nil
}
