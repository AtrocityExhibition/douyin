package douyinComment

import (
	"context"
	"douyin/api/douyin/internal/utils"
	"douyin/rpc/interactive/pb"
	"fmt"

	"douyin/api/douyin/internal/svc"
	"douyin/api/douyin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDouyinCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinCommentListLogic {
	return &DouyinCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DouyinCommentListLogic) DouyinCommentList(req *types.DouyinCommentListRequest) (resp *types.DouyinCommentListResponse, err error) {
	// todo: add your logic here and delete this line
	rpcReq := pb.DouyinCommentListRequest{
		Token:   req.Token,
		VideoId: req.Video_id,
	}
	fmt.Println("调用 rpc DouyinCommentList")
	rpcResp, err := l.svcCtx.InteractiveRpcClient.DouyinCommentList(l.ctx, &rpcReq)
	if err != nil {
		fmt.Println("rpc 服务 DouyinCommentList 出错了")
		return &types.DouyinCommentListResponse{
			Status_code: rpcResp.StatusCode,
			Status_msg:  "rpc 服务 DouyinCommentList 出错了",
		}, err
	}

	var commentList []types.Comment
	for _, c := range rpcResp.CommentList {
		comment := utils.CommentRPC2API(c)
		commentList = append(commentList, comment)
	}

	return &types.DouyinCommentListResponse{
		Status_code:  rpcResp.StatusCode,
		Comment_list: commentList,
	}, nil
}
