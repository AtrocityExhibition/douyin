package logic

import (
	"context"
	"douyin/rpc/social/utils"

	"douyin/rpc/social/internal/svc"
	"douyin/rpc/social/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinRelationFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationFollowListLogic {
	return &DouyinRelationFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinRelationFollowListLogic) DouyinRelationFollowList(in *pb.DouyinRelationFollowListRequest) (*pb.DouyinRelationFollowListResponse, error) {
	// todo: add your logic here and delete this line
	// 从relation中查找user_id == in.user_id
	relationList, err := l.svcCtx.RelationModel.FindByUser(l.ctx, in.UserId)
	if err != nil {
		return &pb.DouyinRelationFollowListResponse{
			StatusCode: 1,
		}, err
	}
	var resToUserList []*pb.User
	for _, relationItem := range relationList {
		toUser, err := l.svcCtx.UserModel.FindOne(l.ctx, relationItem.ToUserId.Int64)
		if err != nil {
			return &pb.DouyinRelationFollowListResponse{
				StatusCode: 1,
				StatusMsg:  "查找关注用户出错",
			}, err
		}
		resToUser, _ := utils.UserModelPb(toUser)

		resToUserList = append(resToUserList, resToUser)

	}
	return &pb.DouyinRelationFollowListResponse{
		StatusCode: 0,
		UserList:   resToUserList,
	}, nil
}
