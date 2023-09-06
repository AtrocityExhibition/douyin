package douyinComment

import (
	"context"
	"douyin/rpc/interactive/pb"
	"fmt"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"
	"douyin/api/douyin/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinCommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinCommentActionLogic {
	return &DouyinCommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinCommentActionLogic) DouyinCommentAction(req *types.DouyinCommentActionRequest) (resp *types.DouyinCommentActionResponse, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinCommentActionRequest{
		Token:       req.Token,
		VideoId:     req.Video_id,
		ActionType:  req.Action_type,
		CommentText: req.Comment_text,
		CommentId:   req.Comment_id,
	}
	fmt.Println("调用 rpc DouyinCommentAction")
	rpcResp, err := l.svcCtx.InteractiveRpcClient.DouyinCommentAction(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("rpc 服务 DouyinCommentAction 出错了", rpcResp.StatusMsg, err)
		return &types.DouyinCommentActionResponse{
			Status_code: rpcResp.StatusCode,
			Status_msg:  "rpc 服务 DouyinCommentAction 出错了",
		}, err
	}

	if req.Action_type == 1 {
		return &types.DouyinCommentActionResponse{
			Status_code: rpcResp.StatusCode,
			Status_msg:  rpcResp.StatusMsg,
			Respond_comment: types.Comment{
				Id:          rpcResp.Comment.Id,
				User:        *utils.UserRPC2API(rpcResp.Comment.User),
				Content:     rpcReq.CommentText,
				Create_date: rpcResp.Comment.CreateDate,
			},
		}, nil
	} else {
		return &types.DouyinCommentActionResponse{
			Status_code: rpcResp.StatusCode,
		}, nil
	}
}
