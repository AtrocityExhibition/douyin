package logic

import (
	"context"
	"douyin/rpc/social/utils"
	"fmt"

	"douyin/rpc/social/internal/svc"
	"douyin/rpc/social/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinRelationFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationFriendListLogic {
	return &DouyinRelationFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinRelationFriendListLogic) DouyinRelationFriendList(in *pb.DouyinRelationFriendListRequest) (*pb.DouyinRelationFriendListResponse, error) {
	// todo: add your logic here and delete this line
	// 从friend表中获取好友表项
	friendList, err := l.svcCtx.FriendModel.FindByUser(l.ctx, in.UserId)
	if err != nil {
		fmt.Println("查找好友列表出错")
		return &pb.DouyinRelationFriendListResponse{
			StatusCode: 1,
			StatusMsg:  "查找好友列表出错",
		}, err
	}
	var resUserList []*pb.User
	for _, friendItem := range friendList {
		// 获取好友的user_id
		friendId := friendItem.User1Id.Int64
		if friendItem.User1Id.Int64 == in.UserId {
			friendId = friendItem.User2Id.Int64
		}
		muser, err := l.svcCtx.UserModel.FindOne(l.ctx, friendId)
		if err != nil {
			fmt.Println("查找好友出错")
			return &pb.DouyinRelationFriendListResponse{
				StatusCode: 1,
				StatusMsg:  "查找好友出错",
			}, err
		}
		user, _ := utils.UserModelPb(muser)
		resUserList = append(resUserList, user)
	}
	return &pb.DouyinRelationFriendListResponse{
		StatusCode: 0,
		UserList:   resUserList,
	}, nil
}
