package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/luan-nguyen-huu/Adam/internal/domain"
	"github.com/luan-nguyen-huu/Adam/internal/utils"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"
)

type UserService struct {
	userRepo domain.UserRepository
	jwtMaker jwt.JWTMakerInterface
	hashCost int
}

func NewUserService(userRepo domain.UserRepository, jwtMaker jwt.JWTMakerInterface, hashCost int) domain.UserService {
	return &UserService{
		userRepo: userRepo,
		jwtMaker: jwtMaker,
		hashCost: hashCost,
	}
}

func (s *UserService) Register(ctx context.Context, name, email, password string) (*domain.AuthTokens, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(ctx, email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(password, s.hashCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:    email,
		Password: hashedPassword,
		Name:     name,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	accessToken, err := s.jwtMaker.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtMaker.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &domain.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (*domain.AuthTokens, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, err
	}

	if err := utils.CheckPasswordHash(password, user.Password); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	accessToken, err := s.jwtMaker.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtMaker.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &domain.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) GetMe(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return s.userRepo.GetByID(ctx, userID)
}

func (s *UserService) RefreshToken(ctx context.Context, userID uuid.UUID) (*domain.AuthTokens, error) {
	// Verify that user still exists in database
	_, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.jwtMaker.GenerateAccessToken(userID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtMaker.GenerateRefreshToken(userID)
	if err != nil {
		return nil, err
	}

	return &domain.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
