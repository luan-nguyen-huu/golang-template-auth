package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTMaker struct {
	secretKey_access string
	secretKey_refresh string
	TLL_access time.Duration
	TLL_refresh time.Duration
}

type JWTMakerInterface interface {
	GenerateAccessToken(userID uuid.UUID) (string, error)
	GenerateRefreshToken(userID uuid.UUID) (string, error)
	VerifyAccessToken(tokenStr string) (*UserClaims, error)
	VerifyRefreshToken(tokenStr string) (*UserClaims, error)
}

func NewJWTMaker(secretKey_access string, secretKey_refresh string, TLL_access time.Duration, TLL_refresh time.Duration) *JWTMaker {
	return &JWTMaker{
		secretKey_access: secretKey_access,
		secretKey_refresh: secretKey_refresh,
		TLL_access: TLL_access,
		TLL_refresh: TLL_refresh,
	}
}

func (maker *JWTMaker) GenerateAccessToken(userID uuid.UUID) (string, error) {
	claims, err := NewUserClaims(userID, maker.TLL_access)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(maker.secretKey_access))
}

func (maker *JWTMaker) GenerateRefreshToken(userID uuid.UUID) (string, error) {
	claims, err := NewUserClaims(userID, maker.TLL_refresh)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(maker.secretKey_refresh))
}

func (maker *JWTMaker) VerifyAccessToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(maker.secretKey_access), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, jwt.ErrInvalidKeyType
	}

	return claims, nil
}

func (maker *JWTMaker) VerifyRefreshToken(tokenStr string) (*UserClaims, error)	 {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(maker.secretKey_refresh), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, jwt.ErrInvalidKeyType
	}

	return claims, nil
}