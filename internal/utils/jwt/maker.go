package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/luan-nguyen-huu/Adam/internal/domain"
)

type JWTMakerInterface interface {
	GenerateAccessToken(userID uuid.UUID) (string, error)
	GenerateRefreshToken(userID uuid.UUID) (string, error)
	VerifyAccessToken(tokenStr string) (*UserClaims, error)
	VerifyRefreshToken(tokenStr string) (*UserClaims, error)
}

type JWTMaker struct {
	accessSecret  string
	refreshSecret string
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

func NewJWTMaker(accessSecret, refreshSecret string, accessTTL, refreshTTL time.Duration) *JWTMaker {
	return &JWTMaker{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessTTL:     accessTTL,
		refreshTTL:    refreshTTL,
	}
}

func (maker *JWTMaker) GenerateAccessToken(userID uuid.UUID) (string, error) {
	claims, err := NewUserClaims(userID, maker.accessTTL)
	if err != nil {
		return "", fmt.Errorf("failed to create access token claims: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(maker.accessSecret))
}

func (maker *JWTMaker) GenerateRefreshToken(userID uuid.UUID) (string, error) {
	claims, err := NewUserClaims(userID, maker.refreshTTL)
	if err != nil {
		return "", fmt.Errorf("failed to create refresh token claims: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(maker.refreshSecret))
}

func (maker *JWTMaker) VerifyAccessToken(tokenStr string) (*UserClaims, error) {
	return maker.verifyToken(tokenStr, maker.accessSecret)
}

func (maker *JWTMaker) VerifyRefreshToken(tokenStr string) (*UserClaims, error) {
	return maker.verifyToken(tokenStr, maker.refreshSecret)
}

func (maker *JWTMaker) verifyToken(tokenStr, secret string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, domain.ErrInvalidToken
			}
			return []byte(secret), nil
		},
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, domain.ErrTokenExpired
		}
		return nil, domain.ErrInvalidToken
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, domain.ErrInvalidToken
	}

	return claims, nil
}