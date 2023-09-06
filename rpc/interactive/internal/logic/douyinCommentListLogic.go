package logic

import (
	"context"

	"douyin/rpc/interactive/internal/svc"
	"douyin/rpc/interactive/pb"
	utils "douyin/rpc/interactive/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinCommentListLogic {
	return &DouyinCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinCommentListLogic) DouyinCommentList(in *pb.DouyinCommentListRequest) (*pb.DouyinCommentListResponse, error) {
	// todo: add your logic here and delete this line
	// 根据video_id查询comment中的评论
	commentResult, err := l.svcCtx.CommentModel.FindByVideo(l.ctx, in.VideoId)
	if err != nil {
		return &pb.DouyinCommentListResponse{
			StatusCode: 1,
			StatusMsg:  "根据video_id查询comment中的评论 失败",
		}, nil
	}

	var commentList []*pb.Comment

	for _, mcomment := range commentResult {
		// 根据mcomment中的user_id查找user
		muser, err := l.svcCtx.UserModel.FindOne(l.ctx, mcomment.UserId.Int64)
		if err != nil {
			return &pb.DouyinCommentListResponse{
				StatusCode: 1,
				StatusMsg:  "根据mcomment UserId 查询user 失败",
			}, nil
		}
		user, err := utils.UserModelPb(muser)

		comment := &pb.Comment{
			Id:         mcomment.Id,
			User:       user,
			Content:    mcomment.Contents.String,
			CreateDate: mcomment.CreateAt.String(),
		}
		commentList = append(commentList, comment)
	}
	return &pb.DouyinCommentListResponse{
		StatusCode:  0,
		CommentList: commentList,
	}, nil
}
