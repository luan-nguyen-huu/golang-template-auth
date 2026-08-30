package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/luan-nguyen-huu/Adam/configs"
	"github.com/luan-nguyen-huu/Adam/internal/middlewares"
	v1 "github.com/luan-nguyen-huu/Adam/internal/routers/v1"
)

func NewMainRouter(cfg *configs.Config, v1Router *v1.V1Router) http.Handler {
	r := chi.NewRouter()

	// Global Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(middlewares.CorsMiddleware(cfg.CORS.GetAllowedOrigins()))

	// Health Check / Root info
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"service":"%s","env":"%s","status":"healthy"}`, cfg.App.Name, cfg.App.Env)))
	})

	// API V1 Routes
	r.Route("/api/v1", func(r chi.Router) {
		v1Router.RegisterRoutes(r)
	})

	return r
}
