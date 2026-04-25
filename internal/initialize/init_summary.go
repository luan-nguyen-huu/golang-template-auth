package initialize

import (
	"gorm.io/gorm"
	"log"

	"github.com/luan-nguyen-huu/Adam/configs"
)

type App struct {
	Config *configs.Config
	DB     *gorm.DB
}

func Init() *App {
	cfg, err := configs.Load()
	if err != nil {
		panic(err)
	}

	db, err := InitDatabase()
	if err != nil {
		panic(err)
	}

	log.Println("Database connected successfully")

	return &App{
		Config: cfg,
		DB:     db,
	}
}
