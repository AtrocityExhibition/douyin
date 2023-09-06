package douyinRelation

import (
	"context"
	"douyin/rpc/social/pb"
	"fmt"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinRelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationActionLogic {
	return &DouyinRelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinRelationActionLogic) DouyinRelationAction(req *types.DouyinRelationActionRequest) (resp *types.DouyinRelationActionResponse, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinRelationActionRequest{
		Token:      req.Token,
		ToUserId:   req.To_user_id,
		ActionType: req.Action_type,
	}
	fmt.Println("调用 rpc DouyinRelationAction")
	rpcResp, err := l.svcCtx.SocialRpcClient.DouyinRelationAction(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("rpc 服务 DouyinRelationAction 出错了")
		return &types.DouyinRelationActionResponse{
			Status_code: rpcResp.StatusCode,
			Status_msg:  "rpc 服务 DouyinRelationAction 出错了",
		}, err
	}
	return &types.DouyinRelationActionResponse{
		Status_code: rpcResp.StatusCode,
	}, nil
}
