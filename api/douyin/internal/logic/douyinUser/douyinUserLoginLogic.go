package douyinUser

import (
	"context"
	"douyin/rpc/core/pb"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinUserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinUserLoginLogic {
	return &DouyinUserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinUserLoginLogic) DouyinUserLogin(req *types.DouyinUserLoginReq) (resp *types.DouyinUserLoginResp, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinUserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}
	rpcResp, err := l.svcCtx.CoreRpcClient.DouyinUserLogin(l.ctx, &rpcReq)
	if err != nil {
		return &types.DouyinUserLoginResp{
			User_id:     rpcResp.UserId,
			Token:       rpcResp.Token,
			Status_code: rpcResp.StatusCode,
		}, err
	}
	return &types.DouyinUserLoginResp{
		User_id:     rpcResp.UserId,
		Token:       rpcResp.Token,
		Status_code: rpcResp.StatusCode,
	}, nil
}
