package logic

import (
	"context"
	"douyin/rpc/core/internal/JWT"
	"douyin/rpc/core/internal/svc"
	"douyin/rpc/core/pb"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type DouyinUserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinUserLoginLogic {
	return &DouyinUserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinUserLoginLogic) DouyinUserLogin(in *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {
	// todo: add your logic here and delete this line
	// todo: 未考虑用户名和密码都相同的情况
	result, err := l.svcCtx.UserModel.FindOneByToken(l.ctx, in.Username, in.Password)
	if err != nil {
		fmt.Println("出错了")
		return &pb.DouyinUserLoginResponse{
			StatusCode: 1,
		}, err
	}
	// 找到了这位用户
	// 生成token
	claims := JWT.TokenClaims{
		Username: result.Name,
		Password: result.Password,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("douyin"))

	fmt.Println(tokenString)

	return &pb.DouyinUserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "登陆成功",
		UserId:     result.Id,
		Token:      tokenString,
	}, nil
}
