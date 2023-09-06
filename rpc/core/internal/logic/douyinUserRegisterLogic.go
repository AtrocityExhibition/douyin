package logic

import (
	"context"
	"douyin/rpc/core/pb"
	"fmt"

	"douyin/model/user"
	"douyin/rpc/core/internal/JWT"
	"douyin/rpc/core/internal/svc"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

const MYSQL_KEY_EXITS = 1062

type DouyinUserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDouyinUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DouyinUserRegisterLogic {
	return &DouyinUserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DouyinUserRegisterLogic) DouyinUserRegister(in *pb.DouyinUserRegisterRequest) (*pb.DouyinUserRegisterResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("进入 DouyinUserRegister rpc 服务")
	u := &user.User{
		Name:     in.Username,
		Password: in.Password,
	}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, u)

	//修改, 在出现两个相同用户名的时候返回错误提示
	if err.(*mysql.MySQLError).Number == MYSQL_KEY_EXITS {
		return &pb.DouyinUserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  "用户名已存在",
		}, nil
	}

	if err != nil {
		return &pb.DouyinUserRegisterResponse{
			StatusCode: 1,
		}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return &pb.DouyinUserRegisterResponse{
			StatusCode: 1,
		}, err
	}

	claims := JWT.TokenClaims{
		Username: u.Name,
		Password: u.Password,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("douyin"))

	if err != nil {
		return &pb.DouyinUserRegisterResponse{
			StatusCode: 1,
		}, err
	}

	fmt.Println(tokenString)

	return &pb.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserId:     id,
		Token:      tokenString,
	}, nil
}
