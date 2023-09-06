package douyinUser

import (
	"context"
	"douyin/rpc/core/pb"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinUserLogic {
	return &DouyinUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinUserLogic) DouyinUser(req *types.DouyinUserReq) (resp *types.DouyinUserResp, err error) {
	// todo: add your logic here and delete this line

	rpcReq := pb.DouyinUserRequest{
		UserId: req.User_id,
		Token:  req.Token,
	}

	rpcResp, err := l.svcCtx.CoreRpcClient.GetUserInfo(l.ctx, &rpcReq)
	user := types.User{
		Id:               rpcResp.User.Id,
		Name:             rpcResp.User.Name,
		Follow_count:     rpcResp.User.FollowCount,
		Follower_count:   rpcResp.User.FollowerCount,
		Avatar:           rpcResp.User.Avatar,
		Background_image: rpcResp.User.BackgroundImage,
		Signature:        rpcResp.User.Signature,
		Total_favorited:  rpcResp.User.TotalFavorited,
		Work_count:       rpcResp.User.WorkCount,
		Favorite_count:   rpcResp.User.FavoriteCount,
	}
	if err != nil {
		return &types.DouyinUserResp{
			Status_code: rpcResp.StatusCode,
			User:        user,
		}, err
	}

	return &types.DouyinUserResp{
		Status_code: rpcResp.StatusCode,
		User:        user,
	}, nil
}
