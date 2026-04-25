package initialize

import (
	"log"
	"github.com/luan-nguyen-huu/Adam/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	cfg, err := configs.Load()
	if err != nil {
		panic(err)
	}

	dsn := configs.GetPostgresDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db, nil
}
