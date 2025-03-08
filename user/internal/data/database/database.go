package database

import (
	"time"

	"user/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	DataConfig config.DataConfig
	GormConfig config.GormConfig
)

func InitDB() {
	if DataConfig.Driver == "postgres" {
		InitPgSQLDB()
	}
}

func InitPgSQLDB() {
	var err error

	g := &gorm.Config{
		PrepareStmt: true, // 预编译SQL提升性能
		NowFunc: func() time.Time { // 统一时间管理，使用UTC时间
			return time.Now().UTC()
		},
	}

	DB, err = gorm.Open(postgres.Open(DataConfig.DSN), g)
	if err != nil {
		logx.Error(err)
		return
	}

	sqlDB, err := DB.DB()
	if err != nil {
		logx.Error(err)
		return
	}

	sqlDB.SetMaxIdleConns(GormConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(GormConfig.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(GormConfig.ConnMaxLifetime) * time.Minute)

	if GormConfig.AutoMigrate {
		err = DB.AutoMigrate(
		// &model.User{},
		)
		if err != nil {
			logx.Error("Auto migration failed:", err)
		} else {
			logx.Info("Auto migration completed successfully")
		}
	}

	logx.Info("Database connection established successfully")
}
