package config

import "github.com/zeromicro/go-zero/rest"

type (
	Config struct {
		rest.RestConf
		S3 S3Config `json:"s3"`
	}

	S3Config struct {
		// S3访问凭证
		AccessKey string `json:"accessKey"`
		SecretKey string `json:"secretKey"`

		// S3基本配置
		Region   string `json:"region"`
		Bucket   string `json:"bucket"`
		Endpoint string `json:"endpoint"`

		// 存储路径配置
		BasePath string `json:"basePath,optional"` // 基础路径，如 "markdown/"

		// 安全配置
		UseSSL bool `json:"useSSL,default=true"`

		// Markdown相关配置
		MaxFileSize  int64    `json:"maxFileSize,default=5242880"` // 默认5MB
		AllowedTypes []string `json:"allowedTypes,optional,default='.md,.markdown'"`

		// 缓存和性能配置
		CacheControl string `json:"cacheControl,optional,default=max-age=86400"` // 默认缓存1天
		Timeout      int    `json:"timeout,default=30"`                          // 默认30秒超时
	}
)
