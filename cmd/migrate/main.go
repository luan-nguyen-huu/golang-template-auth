package main

import (
	"log"

	"github.com/luan-nguyen-huu/Adam/configs"
	"github.com/luan-nguyen-huu/Adam/internal/entities"
	"github.com/luan-nguyen-huu/Adam/internal/initialize"
)

func main() {
	cfg, err := configs.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	db, err := initialize.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect database: %v", err)
	}

	log.Println("🔄 Running database migrations...")

	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migration completed successfully")
}