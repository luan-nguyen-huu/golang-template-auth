package services

import (
	"github.com/google/uuid"

	"github.com/luan-nguyen-huu/Adam/internal/entities"
	"github.com/luan-nguyen-huu/Adam/internal/utils"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"
)

type UserService struct {
	userRepo entities.UserRepositoryInterface
	jwtMaker jwt.JWTMakerInterface
}

func NewUserService(userRepo entities.UserRepositoryInterface, jwtMaker jwt.JWTMakerInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
		jwtMaker: jwtMaker,
	}
}

func (s *UserService) RegisterUser(username string, password string, email string) (string, string, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "","", err
	}
	user := &entities.User{
		Email:        email,
		PasswordHash: hashedPassword,
		Name:         username,
	}
	if err := s.userRepo.CreateUser(user); err != nil {
		return "", "", err
	}

	accessToken, err := s.jwtMaker.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.jwtMaker.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "",  err
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) LoginUser(email string, password string) (string, string, error) {
	var user, err = s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}

	if err := utils.CheckPasswordHash(password, user.PasswordHash); err != nil {
		return "", "", err
	}

	accessToken, err := s.jwtMaker.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.jwtMaker.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "",  err
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) GetMe(userID uuid.UUID) (*entities.User, error) {
	return s.userRepo.GetUserByID(userID)
}

func (s *UserService) RefreshToken(userID uuid.UUID) (string, string, error) {
	accessToken, err := s.jwtMaker.GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := s.jwtMaker.GenerateRefreshToken(userID)
	if err != nil {
		return "", "",  err
	}
	return accessToken, refreshToken, nil
}
