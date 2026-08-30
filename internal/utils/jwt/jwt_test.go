package jwt_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/luan-nguyen-huu/Adam/internal/domain"
	"github.com/luan-nguyen-huu/Adam/internal/utils/jwt"
)

func TestJWTMaker_GenerateAndVerify(t *testing.T) {
	accessSecret := "access-secret-12345678901234567890"
	refreshSecret := "refresh-secret-12345678901234567890"
	accessTTL := 15 * time.Minute
	refreshTTL := 24 * time.Hour

	maker := jwt.NewJWTMaker(accessSecret, refreshSecret, accessTTL, refreshTTL)
	userID := uuid.New()

	// 1. Test Access Token
	accessToken, err := maker.GenerateAccessToken(userID)
	if err != nil {
		t.Fatalf("failed to generate access token: %v", err)
	}
	if accessToken == "" {
		t.Fatalf("expected non-empty access token")
	}

	claims, err := maker.VerifyAccessToken(accessToken)
	if err != nil {
		t.Fatalf("failed to verify access token: %v", err)
	}
	if claims.UserID != userID {
		t.Fatalf("expected userID %v, got %v", userID, claims.UserID)
	}

	// 2. Test Refresh Token
	refreshToken, err := maker.GenerateRefreshToken(userID)
	if err != nil {
		t.Fatalf("failed to generate refresh token: %v", err)
	}

	refreshClaims, err := maker.VerifyRefreshToken(refreshToken)
	if err != nil {
		t.Fatalf("failed to verify refresh token: %v", err)
	}
	if refreshClaims.UserID != userID {
		t.Fatalf("expected userID %v, got %v", userID, refreshClaims.UserID)
	}

	// 3. Test Invalid Token
	_, err = maker.VerifyAccessToken("invalid.token.string")
	if err == nil {
		t.Fatalf("expected error verifying invalid token")
	}

	// 4. Test Cross-Secret verification fails (Access token verified by Refresh secret)
	_, err = maker.VerifyRefreshToken(accessToken)
	if err != domain.ErrInvalidToken {
		t.Fatalf("expected ErrInvalidToken when verifying access token with refresh secret, got: %v", err)
	}
}

func TestJWTMaker_ExpiredToken(t *testing.T) {
	accessSecret := "access-secret-12345678901234567890"
	refreshSecret := "refresh-secret-12345678901234567890"
	accessTTL := -1 * time.Minute // Expired immediately
	refreshTTL := -1 * time.Minute

	maker := jwt.NewJWTMaker(accessSecret, refreshSecret, accessTTL, refreshTTL)
	userID := uuid.New()

	token, err := maker.GenerateAccessToken(userID)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	_, err = maker.VerifyAccessToken(token)
	if err != domain.ErrTokenExpired {
		t.Fatalf("expected ErrTokenExpired, got %v", err)
	}
}
