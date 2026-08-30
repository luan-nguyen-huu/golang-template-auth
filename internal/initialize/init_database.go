package initialize

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/luan-nguyen-huu/Adam/configs"
)

func InitDatabase(cfg *configs.Config) (*gorm.DB, error) {
	dsn := configs.GetPostgresDSN(cfg)
	
	var logMode logger.LogLevel
	if cfg.App.Env == "development" {
		logMode = logger.Info
	} else {
		logMode = logger.Warn
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}
