package logic

import (
	"context"
	"database/sql"
	"douyin/rpc/social/internal/JWT"
	"fmt"

	"douyin/model/message"
	"douyin/rpc/social/internal/svc"
	"douyin/rpc/social/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinMessageActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinMessageActionLogic {
	return &DouyinMessageActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinMessageActionLogic) DouyinMessageAction(in *pb.DouyinMessageActionRequest) (*pb.DouyinMessageActionResponse, error) {
	// todo: add your logic here and delete this line
	if err := JWT.JWTAuthToken(l.svcCtx, l.ctx, in.Token); err != nil {
		return &pb.DouyinMessageActionResponse{
			StatusCode: 1,
			StatusMsg:  "token鉴权出错了",
		}, err
	}
	claims, _ := JWT.JWTAuth(in.Token)
	username := claims["Username"].(string)
	password := claims["Password"].(string)
	u, err := l.svcCtx.UserModel.FindOneByToken(l.ctx, username, password)
	if err != nil {
		fmt.Println("查找用户错误")
	}
	// 把message插入表中
	message := message.Message{
		ToUserId:   sql.NullInt64{in.ToUserId, true},
		FromUserId: sql.NullInt64{u.Id, true},
		Content:    sql.NullString{in.Content, true},
	}
	_, err = l.svcCtx.MessageModel.Insert(l.ctx, &message)
	if err != nil {
		return &pb.DouyinMessageActionResponse{
			StatusCode: 1,
			StatusMsg:  "插入message错误",
		}, err
	}
	return &pb.DouyinMessageActionResponse{
		StatusCode: 0,
	}, nil
}
