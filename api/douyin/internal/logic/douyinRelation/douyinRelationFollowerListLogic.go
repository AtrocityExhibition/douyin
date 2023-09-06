package douyinRelation

import (
	"context"
	"douyin/api/douyin/internal/utils"
	"douyin/rpc/social/pb"
	"fmt"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinRelationFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationFollowerListLogic {
	return &DouyinRelationFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinRelationFollowerListLogic) DouyinRelationFollowerList(req *types.DouyinRelationFollowerListRequest) (resp *types.DouyinRelationFollowerListResponse, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinRelationFollowerListRequest{
		Token:  req.Token,
		UserId: req.User_id,
	}
	fmt.Println("调用 rpc DouyinRelationFollowerList")
	rpcResp, err := l.svcCtx.SocialRpcClient.DouyinRelationFollowerList(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("rpc 服务 DouyinRelationFollowerList 出错了")
		return &types.DouyinRelationFollowerListResponse{
			Status_code: rpcResp.StatusCode,
			Status_msg:  "rpc 服务 DouyinRelationFollowerList 出错了",
		}, err
	}
	var userList []types.User
	for _, u := range rpcResp.UserList {
		user := utils.UserRPCS2API(u)
		userList = append(userList, *user)
	}

	return &types.DouyinRelationFollowerListResponse{
		Status_code: rpcResp.StatusCode,
		UserList:    userList,
	}, nil
}
