package douyinMessage

import (
	"context"
	"douyin/rpc/social/pb"
	"fmt"
	"strconv"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinMessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinMessageChatLogic {
	return &DouyinMessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinMessageChatLogic) DouyinMessageChat(req *types.DouyinMessageChatRequest) (resp *types.DouyinMessageChatResponse, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinMessageChatRequest{
		Token:      req.Token,
		ToUserId:   req.To_user_id,
		PreMsgTime: req.Pre_msg_time,
	}
	fmt.Println("调用 rpc DouyinMessageChat")
	rpcResp, err := l.svcCtx.SocialRpcClient.DouyinMessageChat(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("rpc 服务 DouyinMessageChat 出错了")
		return &types.DouyinMessageChatResponse{
			Status_code: rpcResp.StatusCode,
			Status_msg:  "rpc 服务 DouyinMessageChat 出错了",
		}, err
	}
	var messageList []types.Message
	for _, m := range rpcResp.MessageList {
		createTime, err := strconv.ParseInt(m.CreateTime, 10, 64)
		if err != nil {
			fmt.Println("createTime格式错误")
		}
		tm := types.Message{
			Id:           m.Id,
			To_user_id:   m.ToUserId,
			From_user_id: m.FromUserId,
			Content:      m.Content,
			Create_time:  createTime,
		}
		messageList = append(messageList, tm)
		fmt.Println(tm)
	}

	return &types.DouyinMessageChatResponse{
		Status_code:  0,
		Message_list: messageList,
	}, nil
}
