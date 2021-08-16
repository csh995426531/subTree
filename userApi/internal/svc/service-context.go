package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"subTree/userApi/internal/config"
	"subTree/userRpc/userrpcclient"
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
