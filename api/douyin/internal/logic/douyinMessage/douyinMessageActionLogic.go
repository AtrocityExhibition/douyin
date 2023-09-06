package douyinMessage

import (
	"context"
	"douyin/rpc/social/pb"
	"fmt"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinMessageActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinMessageActionLogic {
	return &DouyinMessageActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinMessageActionLogic) DouyinMessageAction(req *types.DouyinMessageActionRequest) (resp *types.DouyinMessageActionResponse, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinMessageActionRequest{
		Token:      req.Token,
		ToUserId:   req.To_user_id,
		ActionType: req.Action_type,
		Content:    req.Content,
	}
	fmt.Println("调用 rpc DouyinMessageAction")
	rpcResp, err := l.svcCtx.SocialRpcClient.DouyinMessageAction(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("rpc 服务 DouyinMessageAction 出错了")
		return &types.DouyinMessageActionResponse{
			Status_code: rpcResp.StatusCode,
			Status_msg:  "rpc 服务 DouyinMessageAction 出错了",
		}, err
	}
	return &types.DouyinMessageActionResponse{
		Status_code: 0,
	}, nil
}
