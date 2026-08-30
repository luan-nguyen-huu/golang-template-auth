package utils_test

import (
	"testing"

	"github.com/luan-nguyen-huu/Adam/internal/domain"
	"github.com/luan-nguyen-huu/Adam/internal/utils"
)

func TestHashPasswordAndCheck(t *testing.T) {
	password := "mySecretPassword123!"

	hashed, err := utils.HashPassword(password, 10)
	if err != nil {
		t.Fatalf("expected no error hashing password, got %v", err)
	}

	if hashed == password {
		t.Fatalf("expected hash to be different from original password")
	}

	// Verify with correct password
	if err := utils.CheckPasswordHash(password, hashed); err != nil {
		t.Fatalf("expected correct password to match hash, got %v", err)
	}

	// Verify with wrong password
	if err := utils.CheckPasswordHash("wrongPassword", hashed); err != domain.ErrInvalidCredentials {
		t.Fatalf("expected ErrInvalidCredentials for wrong password, got %v", err)
	}
}
