package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"goZeroDemo/Users/thy/Downloads/git/search/internal/config"
	"goZeroDemo/Users/thy/Downloads/git/search/internal/middleware"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
