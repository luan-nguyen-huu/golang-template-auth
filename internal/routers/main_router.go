package routers

import (
	"fmt"
	"net/http"
	"time"

	sub "github.com/luan-nguyen-huu/Adam/internal/initialize/sub"
	"github.com/luan-nguyen-huu/Adam/internal/initialize"
	"github.com/luan-nguyen-huu/Adam/internal/middlewares"
	v1 "github.com/luan-nguyen-huu/Adam/internal/routers/v1"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luan-nguyen-huu/Adam/configs"
)

func RegisterMainRoutes(cfg *configs.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Use(middlewares.CorsMiddleware())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello from %s (%s)", cfg.App.Name, cfg.App.Env)))
	})

	db, err := initialize.InitDatabase()
	if err != nil {
		panic(err)
	}

	services := sub.InitServices(db, cfg)

	v1Router := v1.NewV1Router(services)

	r.Route("/api/v1", func(r chi.Router) {
		v1Router.RegisterV1Routes(r)
	})

	return r
}
