package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/luan-nguyen-huu/Adam/configs"
	"github.com/luan-nguyen-huu/Adam/internal/handlers"
	"github.com/luan-nguyen-huu/Adam/internal/initialize"
	"github.com/luan-nguyen-huu/Adam/internal/repositories"
	"github.com/luan-nguyen-huu/Adam/internal/routers"
	v1 "github.com/luan-nguyen-huu/Adam/internal/routers/v1"
	"github.com/luan-nguyen-huu/Adam/internal/services"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"
)

func main() {
	// 1. Load Configurations
	cfg, err := configs.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	// 2. Initialize Database
	db, err := initialize.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL successfully")

	// 3. Initialize Token Maker & Security
	jwtMaker := jwt.NewJWTMaker(
		cfg.JWT.SecretAccess,
		cfg.JWT.SecretRefresh,
		cfg.JWT.AccessTokenExpire,
		cfg.JWT.RefreshTokenExpire,
	)

	// 4. Dependency Injection (Repository -> Service -> Handler -> Router)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo, jwtMaker, cfg.Password.HashCost)
	userHandler := handlers.NewUserHandler(userService, &cfg.JWT)

	userRouter := v1.NewUserRouter(userHandler, jwtMaker)
	v1Router := v1.NewV1Router(userRouter)
	mainRouter := routers.NewMainRouter(cfg, v1Router)

	// 5. HTTP Server Configuration
	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      mainRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 6. Start Server in background goroutine
	go func() {
		log.Printf("🚀 Server is running on http://%s (Env: %s)", addr, cfg.App.Env)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("❌ Server error: %v", err)
		}
	}()

	// 7. Graceful Shutdown listener
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Shutting down server gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}

	// Close database connections
	if sqlDB, err := db.DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Printf("⚠️ Error closing database connection: %v", err)
		} else {
			log.Println("✅ Database connection closed")
		}
	}

	log.Println("👋 Server exited gracefully")
}
