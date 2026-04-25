package v1

import (
	"github.com/luan-nguyen-huu/Adam/internal/entities"
	"github.com/luan-nguyen-huu/Adam/internal/handlers"
	"github.com/luan-nguyen-huu/Adam/internal/middlewares"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserService entities.UserServiceInterface
	TokenMaker  jwt.JWTMakerInterface
}

func NewUserRouter(
	userService entities.UserServiceInterface,
	tokenMaker jwt.JWTMakerInterface,
) *UserRouter {
	return &UserRouter{
		UserService: userService,
		TokenMaker:  tokenMaker,
	}
}

func (ur *UserRouter) RegisterUserRoutes(r chi.Router) {
	userHandler := handlers.NewUserHandler(ur.UserService)

	r.Post("/register", userHandler.RegisterUser)
	r.Post("/login", userHandler.LoginUser)

	r.With(
		middlewares.AuthMiddlewareByCookie(ur.TokenMaker),
	).Get("/me", userHandler.GetMe)

	r.With(
		middlewares.RefreshTokenMiddleware(ur.TokenMaker),
	).Post("/refresh", userHandler.RefreshToken)
}
