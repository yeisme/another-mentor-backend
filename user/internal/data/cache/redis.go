package cache

import (
	"time"
	"user/internal/config"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

var RDC *redis.Client

func NewCache(cacheConfig config.CacheConfig) *redis.Client {
	RDC = InitRedis(cacheConfig)
	if RDC == nil {
		logx.Error("Init redis failed")
		return nil
	}
	logx.Info("Init redis success")
	return RDC
}

func InitRedis(cacheConfig config.CacheConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         cacheConfig.Redis.Addr,
		Password:     cacheConfig.Redis.Password,
		DB:           cacheConfig.Redis.DB,
		PoolSize:     cacheConfig.Redis.PoolSize,
		MinIdleConns: cacheConfig.Redis.MinIdleConns,
		ReadTimeout:  time.Duration(cacheConfig.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cacheConfig.Redis.WriteTimeout) * time.Second,
	})
}
