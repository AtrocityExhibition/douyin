package logic

import (
	"context"
	"database/sql"
	"douyin/model/friend"
	"douyin/model/relation"
	"douyin/rpc/social/internal/JWT"
	"errors"
	"fmt"

	"douyin/rpc/social/internal/svc"
	"douyin/rpc/social/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinRelationActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinRelationActionLogic {
	return &DouyinRelationActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinRelationActionLogic) DouyinRelationAction(in *pb.DouyinRelationActionRequest) (*pb.DouyinRelationActionResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("进入rpc DouyinRelationAction")
	if err := JWT.JWTAuthToken(l.svcCtx, l.ctx, in.Token); err != nil {
		return &pb.DouyinRelationActionResponse{
			StatusCode: 1,
			StatusMsg:  "token鉴权出错了",
		}, err
	}
	// 从token中解析出用户名
	claims, err := JWT.JWTAuth(in.Token)
	if err != nil {
		return &pb.DouyinRelationActionResponse{
			StatusCode: 1,
			StatusMsg:  "token错误",
		}, err
	}
	username := claims["Username"].(string)
	password := claims["Password"].(string)
	switch in.ActionType {
	case 1:
		// 关注
		// 1. 将关注信息插入relation中
		user, err := l.svcCtx.UserModel.FindOneByToken(l.ctx, username, password)
		if err != nil {
			return &pb.DouyinRelationActionResponse{
				StatusCode: 1,
			}, errors.New("查找用户信息失败")
		}
		relationItem := relation.Relation{
			UserId:   sql.NullInt64{Int64: user.Id, Valid: true}, // 从token解析user_id
			ToUserId: sql.NullInt64{Int64: in.ToUserId, Valid: true},
		}
		_, err = l.svcCtx.RelationModel.Insert(l.ctx, &relationItem)
		if err != nil {
			return &pb.DouyinRelationActionResponse{
				StatusCode: 1,
			}, errors.New("插入失败")
		}
		// 2. 修改user的foll_count
		user.FollowCount++
		err = l.svcCtx.UserModel.Update(l.ctx, user)
		if err != nil {
			return &pb.DouyinRelationActionResponse{
				StatusCode: 1,
			}, errors.New("更新user 的 FollowCount 失败")
		}
		// 3. 修改to_user 的 follwer_count
		toUser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.ToUserId)
		if err != nil {
			return &pb.DouyinRelationActionResponse{
				StatusCode: 1,
			}, errors.New("查找被关注的用户to_user失败")
		}
		toUser.FollowerCount++
		err = l.svcCtx.UserModel.Update(l.ctx, toUser)
		if err != nil {
			return &pb.DouyinRelationActionResponse{
				StatusCode: 1,
			}, errors.New("更新to_user 的 FollowerCount 失败")
		}
		// 4. 如果对方也关注了我，就把俩人加入friend中
		result, _ := l.svcCtx.RelationModel.FindOneByUserToUser(l.ctx, in.ToUserId, user.Id)
		if result != nil {
			// 对方也关注了我
			// 插入friend中
			l.svcCtx.FriendModel.Insert(l.ctx, &friend.Friend{
				User1Id: sql.NullInt64{user.Id, true},
				User2Id: sql.NullInt64{in.ToUserId, true},
			})
		}
	case 2:
		// 取消关注
		// 1. 将关注信息从relation中删除
		user, err := l.svcCtx.UserModel.FindOneByToken(l.ctx, username, password)
		err = l.svcCtx.RelationModel.DeleteByUser(l.ctx, user.Id, in.ToUserId)
		if err != nil {
			return &pb.DouyinRelationActionResponse{
				StatusCode: 1,
			}, errors.New("从relation中删除失败")
		}
		// 2. 修改user的foll_count
		user.FollowCount--
		err = l.svcCtx.UserModel.Update(l.ctx, user)
		if err != nil {
			return &pb.DouyinRelationActionResponse{
				StatusCode: 1,
			}, errors.New("更新user 的 FollowCount 失败")
		}
		// 3. 修改to_user 的 follwer_count
		toUser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.ToUserId)
		if err != nil {
			return &pb.DouyinRelationActionResponse{
				StatusCode: 1,
			}, errors.New("查找被关注的用户to_user失败")
		}
		toUser.FollowerCount--
		err = l.svcCtx.UserModel.Update(l.ctx, toUser)
		if err != nil {
			return &pb.DouyinRelationActionResponse{
				StatusCode: 1,
			}, errors.New("更新to_user 的 FollowerCount 失败")
		}
	default:
		return &pb.DouyinRelationActionResponse{
			StatusCode: 1,
		}, errors.New("请输入正确的操作")
	}
	return &pb.DouyinRelationActionResponse{
		StatusCode: 0,
	}, nil
}
