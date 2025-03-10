package config

import "github.com/zeromicro/go-zero/rest"

type (
	Config struct {
		rest.RestConf
		DataConfig DataConfig `json:"data"`
		Auth       Auth       `json:"Auth"`
	}

	DataConfig struct {
		// 数据库驱动
		Driver string `json:"driver"`
		// 数据库连接
		DSN string `json:"dsn"`
		// 关系型数据库配置 TODO: 目前使用gorm
		GormConfig GormConfig `json:"gorm"`
	}

	// GORM配置
	GormConfig struct {
		// 最大闲置连接数
		MaxIdleConns int `json:"maxIdleConns"`
		// 最大开放连接数
		MaxOpenConns int `json:"maxOpenConns"`
		// 连接最大生命周期
		ConnMaxLifetime int `json:"connMaxLifetime"`
		// 是否启用自动迁移
		AutoMigrate bool `json:"autoMigrate"`
	}

	// Auth 鉴权配置
	Auth struct {
		// 访问密钥
		AccessSecret string `json:"AccessSecret"`
		// 访问密钥过期时间
		AccessExpire int `json:"AccessExpire"`
	}
)
