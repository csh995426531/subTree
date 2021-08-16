package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"subModule/userApi/internal/config"
	"subModule/userRpc/userrpcclient"
)

type ServiceContext struct {
	Config config.Config
	UserRpc userrpcclient.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpc: userrpcclient.NewUserRpc(zrpc.MustNewClient(c.UserRpc)),
	}
}
