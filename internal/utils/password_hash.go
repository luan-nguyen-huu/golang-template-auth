package utils

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/luan-nguyen-huu/Adam/internal/domain"
)

// HashPassword hashes a plain password using bcrypt with specified cost (fallback to DefaultCost if <= 0).
func HashPassword(password string, cost int) (string, error) {
	if cost <= 0 {
		cost = bcrypt.DefaultCost
	}
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPasswordHash compares password with hashed password and returns domain.ErrInvalidCredentials on mismatch.
func CheckPasswordHash(password, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return domain.ErrInvalidCredentials
	}
	return nil
}
