package data

import (
	"markdown/internal/config"
	"markdown/internal/data/database"
	"markdown/internal/data/oss"
)

func InitDataService(c config.Config) {
	database.InitDB(c.DataConfig)
	oss.InitOss(c.S3Config)
}
