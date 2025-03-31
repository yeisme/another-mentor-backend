package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"markdown/internal/config"
	"markdown/internal/middleware"
)

type ServiceContext struct {
	Config   config.Config
	UserAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserAuth: middleware.NewUserAuthMiddleware().Handle,
	}
}
