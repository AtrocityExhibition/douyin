package svc

import (
	"douyin/model/comment"
	"douyin/model/favorite"
	"douyin/model/user"
	"douyin/model/video"
	"douyin/rpc/interactive/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     user.UserModel
	VideoModel    video.VideoModel
	FavoriteModel favorite.FavoriteModel
	CommentModel  comment.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     user.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		VideoModel:    video.NewVideoModel(sqlx.NewMysql(c.DB.DataSource)),
		FavoriteModel: favorite.NewFavoriteModel(sqlx.NewMysql(c.DB.DataSource)),
		CommentModel:  comment.NewCommentModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
