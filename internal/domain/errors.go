package domain

import "errors"

// Sentinel Domain Errors
var (
	// User & Auth errors
	ErrUserNotFound          = errors.New("user not found")
	ErrUserAlreadyExists     = errors.New("user with this email already exists")
	ErrInvalidCredentials    = errors.New("incorrect email or password")
	ErrInvalidToken          = errors.New("invalid or malformed authorization token")
	ErrTokenExpired          = errors.New("authorization token has expired")
	ErrMissingAuthToken      = errors.New("missing authorization token")
	ErrInvalidAuthHeader     = errors.New("invalid authorization header format")

	// Common request & server errors
	ErrBadRequest            = errors.New("invalid request payload")
	ErrValidationFailed      = errors.New("validation failed")
	ErrInternalServer        = errors.New("internal server error")
)
