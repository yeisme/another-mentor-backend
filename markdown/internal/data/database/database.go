package database

import (
	"markdown/internal/config"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dataConfig config.DataConfig) {
	logx.Infof("Initializing %s server connection", dataConfig.Driver)
	if dataConfig.Driver == "postgres" {
		InitPgSQLDB(dataConfig)
	}

	logx.Infof("%s server connection established successfully", dataConfig.Driver)
}

func InitPgSQLDB(dataConfig config.DataConfig) {
	var err error

	g := &gorm.Config{
		PrepareStmt: true, // 预编译SQL提升性能
		NowFunc: func() time.Time { // 统一时间管理，使用UTC时间
			return time.Now().UTC()
		},
	}

	DB, err = gorm.Open(postgres.Open(dataConfig.DSN), g)
	if err != nil {
		logx.Errorf("数据库连接失败: %v", err)
		return
	} else {
		logx.Infof("数据库连接成功")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		logx.Errorf("获取数据库连接失败: %v", err)
		return
	}

	sqlDB.SetMaxIdleConns(dataConfig.GormConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dataConfig.GormConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(dataConfig.GormConfig.ConnMaxLifetime) * time.Minute)

	if dataConfig.GormConfig.AutoMigrate {
		err = DB.AutoMigrate()
		if err != nil {
			logx.Errorf("Auto migration failed: %v", err)
		} else {
			logx.Infof("Auto migration completed successfully")
		}
	}

	logx.Info("PostgreSQL Database connection established successfully")
}
