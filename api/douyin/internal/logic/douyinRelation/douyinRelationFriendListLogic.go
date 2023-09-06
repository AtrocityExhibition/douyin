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

type DouyinRelationFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationFriendListLogic {
	return &DouyinRelationFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinRelationFriendListLogic) DouyinRelationFriendList(req *types.DouyinRelationFriendListRequest) (resp *types.DouyinRelationFriendListResponse, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinRelationFriendListRequest{
		Token:  req.Token,
		UserId: req.User_id,
	}
	fmt.Println("调用 rpc DouyinRelationFollowerList")
	rpcResp, err := l.svcCtx.SocialRpcClient.DouyinRelationFriendList(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("rpc 服务 DouyinRelationFollowerList 出错了")
		return &types.DouyinRelationFriendListResponse{
			Status_code: rpcResp.StatusCode,
			Status_msg:  "rpc 服务 DouyinRelationFollowerList 出错了",
		}, err
	}
	var userList []types.User
	for _, u := range rpcResp.UserList {
		user := utils.UserRPCS2API(u)
		userList = append(userList, *user)
	}

	return &types.DouyinRelationFriendListResponse{
		Status_code: rpcResp.StatusCode,
		UserList:    userList,
	}, nil
}
