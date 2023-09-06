package douyinFavorite

import (
	"context"
	"douyin/rpc/interactive/pb"
	"fmt"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinFavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinFavoriteActionLogic {
	return &DouyinFavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinFavoriteActionLogic) DouyinFavoriteAction(req *types.DouyinFavoriteActionRequest) (resp *types.DouyinFavoriteActionResponse, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinFavoriteActionRequest{
		Token:      req.Token,
		VideoId:    req.Video_id,
		ActionType: req.Action_type,
	}
	fmt.Println("调用 rpc DouyinFavoriteAction")
	rpcResp, err := l.svcCtx.InteractiveRpcClient.DouyinFavoriteAction(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("rpc 服务 DouyinFavoriteAction 出错了")
		return &types.DouyinFavoriteActionResponse{
			Status_code: rpcResp.StatusCode,
			Status_msg:  "rpc 服务 DouyinFavoriteAction 出错了",
		}, err
	}
	return &types.DouyinFavoriteActionResponse{
		Status_code: rpcResp.StatusCode,
	}, nil
}
