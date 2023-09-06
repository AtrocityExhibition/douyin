// Code generated by goctl. DO NOT EDIT.
// Source: interactive.proto

package server

import (
	"context"

	"douyin/rpc/interactive/internal/logic"
	"douyin/rpc/interactive/internal/svc"
	"douyin/rpc/interactive/pb"
)

type InteractiveServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedInteractiveServer
}

func NewInteractiveServer(svcCtx *svc.ServiceContext) *InteractiveServer {
	return &InteractiveServer{
		svcCtx: svcCtx,
	}
}

func (s *InteractiveServer) DouyinFavoriteAction(ctx context.Context, in *pb.DouyinFavoriteActionRequest) (*pb.DouyinFavoriteActionResponse, error) {
	l := logic.NewDouyinFavoriteActionLogic(ctx, s.svcCtx)
	return l.DouyinFavoriteAction(in)
}

func (s *InteractiveServer) DouyinFavoriteList(ctx context.Context, in *pb.DouyinFavoriteListRequest) (*pb.DouyinFavoriteListResponse, error) {
	l := logic.NewDouyinFavoriteListLogic(ctx, s.svcCtx)
	return l.DouyinFavoriteList(in)
}

func (s *InteractiveServer) DouyinCommentAction(ctx context.Context, in *pb.DouyinCommentActionRequest) (*pb.DouyinCommentActionResponse, error) {
	l := logic.NewDouyinCommentActionLogic(ctx, s.svcCtx)
	return l.DouyinCommentAction(in)
}

func (s *InteractiveServer) DouyinCommentList(ctx context.Context, in *pb.DouyinCommentListRequest) (*pb.DouyinCommentListResponse, error) {
	l := logic.NewDouyinCommentListLogic(ctx, s.svcCtx)
	return l.DouyinCommentList(in)
}
