package config

import "github.com/zeromicro/go-zero/rest"

type (
	Config struct {
		rest.RestConf
	}

	S3Config struct {
		AccessKey  string `json:"accessKe"`
		SecretKey  string `json:"secretKey"`
		Region     string `json:"region"`
		Bucket     string `json:"bucket"`
		Endpoint   string `json:"endpoint"`
		BasePath   string `json:"basePath"`
		UseSSL     bool   `json:"useSSL"`
		MaxFileSize int64  `json:"maxFileSize"`
		AllowedTypes []string `json:"allowedTypes"`
		CacheControl string `json:"cacheControl"`
		Timeout      int    `json:"timeout"`
	}
)
