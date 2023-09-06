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

type DouyinRelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationFollowListLogic {
	return &DouyinRelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinRelationFollowListLogic) DouyinRelationFollowList(req *types.DouyinRelationFollowListRequest) (resp *types.DouyinRelationFollowListResponse, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinRelationFollowListRequest{
		Token:  req.Token,
		UserId: req.User_id,
	}
	fmt.Println("调用 rpc DouyinRelationFollowList")
	rpcResp, err := l.svcCtx.SocialRpcClient.DouyinRelationFollowList(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("rpc 服务 DouyinRelationFollowList 出错了")
		return &types.DouyinRelationFollowListResponse{
			Status_code: rpcResp.StatusCode,
			Status_msg:  "rpc 服务 DouyinRelationFollowList 出错了",
		}, err
	}
	var userList []types.User
	for _, u := range rpcResp.UserList {
		user := utils.UserRPCS2API(u)
		userList = append(userList, *user)
	}

	return &types.DouyinRelationFollowListResponse{
		Status_code: rpcResp.StatusCode,
		UserList:    userList,
	}, nil
}
