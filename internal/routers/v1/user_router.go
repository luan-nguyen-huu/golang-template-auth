package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/luan-nguyen-huu/Adam/internal/handlers"
	"github.com/luan-nguyen-huu/Adam/internal/middlewares"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"
)

type UserRouter struct {
	userHandler *handlers.UserHandler
	tokenMaker  jwt.JWTMakerInterface
}

func NewUserRouter(userHandler *handlers.UserHandler, tokenMaker jwt.JWTMakerInterface) *UserRouter {
	return &UserRouter{
		userHandler: userHandler,
		tokenMaker:  tokenMaker,
	}
}

func (ur *UserRouter) RegisterRoutes(r chi.Router) {
	// Public routes
	r.Post("/register", ur.userHandler.Register)
	r.Post("/login", ur.userHandler.Login)
	r.Post("/logout", ur.userHandler.Logout)

	// Protected routes
	r.With(middlewares.AuthMiddleware(ur.tokenMaker)).Get("/me", ur.userHandler.GetMe)
	r.With(middlewares.RefreshTokenMiddleware(ur.tokenMaker)).Post("/refresh", ur.userHandler.RefreshToken)
}
