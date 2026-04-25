package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/luan-nguyen-huu/Adam/internal/initialize/sub"
)

type V1Router struct {
	Services *sub.Services
}

func NewV1Router(svcs *sub.Services) *V1Router {
	return &V1Router{
		Services: svcs,
	}
}

func (vr *V1Router) RegisterV1Routes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		userRouter := NewUserRouter(
			vr.Services.UserService,
			vr.Services.TokenMaker,
		)
		userRouter.RegisterUserRoutes(r)
	})
}