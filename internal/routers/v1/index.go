package v1

import (
	"github.com/go-chi/chi/v5"
)

type V1Router struct {
	userRouter *UserRouter
}

func NewV1Router(userRouter *UserRouter) *V1Router {
	return &V1Router{
		userRouter: userRouter,
	}
}

func (vr *V1Router) RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		vr.userRouter.RegisterRoutes(r)
	})
}