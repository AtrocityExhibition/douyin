package logic

import (
	"context"
	"douyin/rpc/core/internal/JWT"
	"douyin/rpc/core/internal/svc"
	"douyin/rpc/core/pb"
	"douyin/rpc/core/utils"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.DouyinUserRequest) (*pb.DouyinUserResponse, error) {
	// todo: add your logic here and delete this line
	if err := JWT.JWTAuthToken(l.svcCtx, l.ctx, in.Token); err != nil {
		return &pb.DouyinUserResponse{
			StatusCode: 1,
			StatusMsg:  "token鉴权出错了",
		}, err
	}
	result, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return &pb.DouyinUserResponse{
			StatusCode: 1,
			StatusMsg:  "根据 user_id 查找用户出错了",
		}, err
	}
	user, _ := utils.UserModelPb(result)

	return &pb.DouyinUserResponse{
		StatusCode: 0,
		User:       user,
	}, nil
}
