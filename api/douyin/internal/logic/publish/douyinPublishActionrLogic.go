package publish

import (
	"context"
	"douyin/rpc/core/pb"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinPublishActionrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinPublishActionrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinPublishActionrLogic {
	return &DouyinPublishActionrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinPublishActionrLogic) DouyinPublishActionr(req *types.DouyinPublishActionrReq) (resp *types.DouyinPublishActionResp, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinPublishActionRequest{
		Token: req.Token,
		Data:  req.Data,
		Title: req.Title,
	}

	rpcResp, err := l.svcCtx.CoreRpcClient.DouyinPublishAction(l.ctx, &rpcReq)
	if err != nil {
		return &types.DouyinPublishActionResp{
			Status_code: rpcResp.StatusCode,
		}, err
	}

	return &types.DouyinPublishActionResp{
		Status_code: rpcResp.StatusCode,
	}, nil
}
