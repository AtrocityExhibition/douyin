package svc

import (
	"douyin/model/favorite"
	"douyin/model/relation"
	"douyin/model/user"
	"douyin/model/video"
	"douyin/rpc/core/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     user.UserModel
	VideoModel    video.VideoModel
	Relationmodel relation.RelationModel
	FavoriteModel favorite.FavoriteModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     user.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		VideoModel:    video.NewVideoModel(sqlx.NewMysql(c.DB.DataSource)),
		FavoriteModel: favorite.NewFavoriteModel(sqlx.NewMysql(c.DB.DataSource)),
		Relationmodel: relation.NewRelationModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
