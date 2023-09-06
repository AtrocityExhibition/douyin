package douyinUser

import (
	"context"
	"douyin/api/core/internal/svc"
	"douyin/api/core/internal/types"
	"douyin/rpc/core/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinUserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinUserRegisterLogic {
	return &DouyinUserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinUserRegisterLogic) DouyinUserRegister(req *types.DouyinUserRegisterReq) (resp *types.DouyinUserRegisterResp, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}
	rpcResp, err := l.svcCtx.CoreRpcClient.DouyinUserRegister(l.ctx, &rpcReq)
	if err != nil {
		return &types.DouyinUserRegisterResp{
			User_id:     rpcResp.UserId,
			Token:       rpcResp.Token,
			Status_code: rpcResp.StatusCode,
		}, err
	}
	return &types.DouyinUserRegisterResp{
		User_id:     rpcResp.UserId,
		Token:       rpcResp.Token,
		Status_code: rpcResp.StatusCode,
	}, nil
}
