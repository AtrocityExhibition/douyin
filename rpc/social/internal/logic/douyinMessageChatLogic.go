package logic

import (
	"context"
	"douyin/rpc/social/internal/JWT"
	"fmt"

	"douyin/rpc/social/internal/svc"
	"douyin/rpc/social/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinMessageChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinMessageChatLogic {
	return &DouyinMessageChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinMessageChatLogic) DouyinMessageChat(in *pb.DouyinMessageChatRequest) (*pb.DouyinMessageChatResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("进入 rpc DouyinMessageChat")
	if err := JWT.JWTAuthToken(l.svcCtx, l.ctx, in.Token); err != nil {
		fmt.Println("token鉴权出错了")
		return &pb.DouyinMessageChatResponse{
			StatusCode: 1,
			StatusMsg:  "token鉴权出错了",
		}, err
	}
	claims, _ := JWT.JWTAuth(in.Token)
	username := claims["Username"].(string)
	password := claims["Password"].(string)
	u, err := l.svcCtx.UserModel.FindOneByToken(l.ctx, username, password)
	if err != nil {
		return &pb.DouyinMessageChatResponse{
			StatusCode: 1,
			StatusMsg:  "查找用户错误",
		}, err
	}
	messages, err := l.svcCtx.MessageModel.FindByUserToUser(l.ctx, u.Id, in.ToUserId)
	if err != nil {
		fmt.Println("查找message中的聊天记录错误")
		return &pb.DouyinMessageChatResponse{
			StatusCode: 1,
			StatusMsg:  "查找message中的聊天记录错误",
		}, err
	}
	fmt.Println(messages)
	var messageList []*pb.Message
	for _, m := range messages {
		pm := pb.Message{
			Id:         m.Id,
			ToUserId:   m.ToUserId.Int64,
			FromUserId: m.FromUserId.Int64,
			Content:    m.Content.String,
			CreateTime: m.CreateAt.String(),
		}
		messageList = append(messageList, &pm)
	}
	return &pb.DouyinMessageChatResponse{
		StatusCode:  0,
		MessageList: messageList,
	}, nil
}
