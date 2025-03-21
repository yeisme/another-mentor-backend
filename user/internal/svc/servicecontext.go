package svc

import (
	"user/internal/config"
	"user/internal/data/cache"
	"user/internal/data/repository"
	"user/internal/middleware"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	AdminAuth rest.Middleware
	UserAuth  rest.Middleware
	// 添加Repository字段
    UserRepo repository.UserRepository
	AdminRepo repository.AdminRepository
	// Cache
	Redis *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		AdminAuth: middleware.NewAdminAuthMiddleware(c).Handle,
		UserAuth:  middleware.NewUserAuthMiddleware(c).Handle,
		UserRepo:  repository.NewUserRepository(),
		AdminRepo: repository.NewAdminRepository(),
		Redis:     cache.NewCache(c.DataConfig.CacheConfig),
	}
}
