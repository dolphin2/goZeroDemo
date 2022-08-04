package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"goZeroDemo/Users/thy/Downloads/git/order/api/internal/config"
	"goZeroDemo/Users/thy/Downloads/git/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}