package svc

import (
	"user/internal/config"
	"user/internal/data/repository"
	"user/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	AdminAuth rest.Middleware
	UserAuth  rest.Middleware
	// 添加Repository字段
    UserRepo repository.UserRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		AdminAuth: middleware.NewAdminAuthMiddleware(c).Handle,
		UserAuth:  middleware.NewUserAuthMiddleware(c).Handle,
		UserRepo: repository.NewUserRepository(),

	}
}
