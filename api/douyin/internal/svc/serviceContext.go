package svc

import (
	"douyin/api/douyin/internal/config"
	"douyin/rpc/core/core"
	"douyin/rpc/interactive/interactive"
	"douyin/rpc/social/social"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	CoreRpcClient        core.Core
	InteractiveRpcClient interactive.Interactive
	SocialRpcClient      social.Social
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		CoreRpcClient:        core.NewCore(zrpc.MustNewClient(c.CoreRpcConf)),
		InteractiveRpcClient: interactive.NewInteractive(zrpc.MustNewClient(c.InteractiveRpcConf)),
		SocialRpcClient:      social.NewSocial(zrpc.MustNewClient(c.SocialRpcConf)),
	}
}
