package sub

import (
	"gorm.io/gorm"

	"github.com/luan-nguyen-huu/Adam/configs"
	"github.com/luan-nguyen-huu/Adam/internal/entities"
	"github.com/luan-nguyen-huu/Adam/internal/repositories"
	"github.com/luan-nguyen-huu/Adam/internal/services"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"
)

type Services struct {
	UserService entities.UserServiceInterface
	TokenMaker  jwt.JWTMakerInterface
}

func InitServices(db *gorm.DB, cfg *configs.Config) *Services {
	jwtMaker := jwt.NewJWTMaker(
		cfg.JWT.SecretAccess,
		cfg.JWT.SecretRefresh,
		cfg.JWT.AccessTokenExpire,
		cfg.JWT.RefreshTokenExpire,
	)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo, jwtMaker)

	return &Services{
		UserService: userService,
		TokenMaker:  jwtMaker,
	}
}