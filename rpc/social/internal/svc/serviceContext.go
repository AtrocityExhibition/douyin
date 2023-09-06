package svc

import (
	"douyin/model/friend"
	"douyin/model/message"
	"douyin/model/relation"
	"douyin/model/user"
	"douyin/rpc/social/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     user.UserModel
	RelationModel relation.RelationModel
	FriendModel   friend.FriendModel
	MessageModel  message.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     user.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		RelationModel: relation.NewRelationModel(sqlx.NewMysql(c.DB.DataSource)),
		FriendModel:   friend.NewFriendModel(sqlx.NewMysql(c.DB.DataSource)),
		MessageModel:  message.NewMessageModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
