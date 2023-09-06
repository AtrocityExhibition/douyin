package logic

import (
	"context"
	"douyin/rpc/social/utils"

	"douyin/rpc/social/internal/svc"
	"douyin/rpc/social/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinRelationFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationFollowerListLogic {
	return &DouyinRelationFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinRelationFollowerListLogic) DouyinRelationFollowerList(in *pb.DouyinRelationFollowerListRequest) (*pb.DouyinRelationFollowerListResponse, error) {
	// todo: add your logic here and delete this line
	// 从relation中查找to_user_id == in.ser_id
	relationList, err := l.svcCtx.RelationModel.FindByToUser(l.ctx, in.UserId)
	if err != nil {
		return &pb.DouyinRelationFollowerListResponse{
			StatusCode: 1,
		}, err
	}
	var resFollowerList []*pb.User
	for _, relationItem := range relationList {
		follower, err := l.svcCtx.UserModel.FindOne(l.ctx, relationItem.UserId.Int64)
		if err != nil {
			return &pb.DouyinRelationFollowerListResponse{
				StatusCode: 1,
				StatusMsg:  "查找关注用户出错",
			}, err
		}
		resFollower, _ := utils.UserModelPb(follower)

		resFollowerList = append(resFollowerList, resFollower)

	}
	return &pb.DouyinRelationFollowerListResponse{
		StatusCode: 0,
		UserList:   resFollowerList,
	}, nil
}
