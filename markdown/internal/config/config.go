package config

import "github.com/zeromicro/go-zero/rest"

type (
	Config struct {
		rest.RestConf
		S3Config   S3Config   `json:"s3"`
		DataConfig DataConfig `json:"data"`
	}

	S3Config struct {
		AccessKey    string   `json:"accessKey"`
		SecretKey    string   `json:"secretKey"`
		Region       string   `json:"region"`
		Bucket       string   `json:"bucket"`
		Endpoint     string   `json:"endpoint"`
		BasePath     string   `json:"basePath"`
		UseSSL       bool     `json:"useSSL"`
		MaxFileSize  int64    `json:"maxFileSize"`
		AllowedTypes []string `json:"allowedTypes"`
		CacheControl string   `json:"cacheControl"`
		Timeout      int      `json:"timeout"`
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
)
