package logic

import (
	"context"
	"database/sql"
	"douyin/model/comment"
	"errors"

	"douyin/rpc/interactive/internal/svc"
	"douyin/rpc/interactive/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinCommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinCommentActionLogic {
	return &DouyinCommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinCommentActionLogic) DouyinCommentAction(in *pb.DouyinCommentActionRequest) (*pb.DouyinCommentActionResponse, error) {
	// todo: add your logic here and delete this line
	if in.ActionType == 1 {
		// 添加评论
		// 1.video comment_count ++
		video, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
		if err != nil {
			return &pb.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  "1.查询video失败",
			}, err
		}
		video.CommentCount++
		err = l.svcCtx.VideoModel.Update(l.ctx, video)
		if err != nil {
			return &pb.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  "2.更新video CommentCount失败",
			}, err
		}
		// 2.插入comment user_id video_id contents
		commentStr := comment.Comment{
			UserId:   sql.NullInt64{1, true}, // 从token解析
			VideoId:  sql.NullInt64{in.VideoId, true},
			Contents: sql.NullString{in.CommentText, true},
		}
		_, err = l.svcCtx.CommentModel.Insert(l.ctx, &commentStr)
		if err != nil {
			return &pb.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  "3.插入 comment 失败",
			}, err
		}
	} else if in.ActionType == 2 {
		// 删除评论
		// 1.video comment_count --
		video, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
		if err != nil {
			return &pb.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  "1.查询video失败",
			}, err
		}
		video.CommentCount--
		err = l.svcCtx.VideoModel.Update(l.ctx, video)
		if err != nil {
			return &pb.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  "2.更新video CommentCount失败",
			}, err
		}
		// 2.删除comment user_id video_id contents
		// 根据user_id video_id 查找评论
		// 从token中解析user
		// 查找user_id
		commentItem, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.CommentId)
		if err != nil {
			return &pb.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  "3.查找 comment 失败",
			}, err
		}
		err = l.svcCtx.CommentModel.Delete(l.ctx, commentItem.Id)
		if err != nil {
			return &pb.DouyinCommentActionResponse{
				StatusCode: 1,
				StatusMsg:  "4.删除 comment 失败",
			}, err
		}
	} else {
		return &pb.DouyinCommentActionResponse{
			StatusCode: 1,
			StatusMsg:  "action type 错误",
		}, errors.New("请输入正确的操作")
	}

	return &pb.DouyinCommentActionResponse{
		StatusCode: 0,
	}, nil
}
