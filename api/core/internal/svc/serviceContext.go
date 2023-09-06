package svc

import (
	"douyin/api/core/internal/config"
	"douyin/rpc/core/core"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	CoreRpcClient core.Core
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CoreRpcClient: core.NewCore(zrpc.MustNewClient(c.CoreRpcConf)),
	}
}
