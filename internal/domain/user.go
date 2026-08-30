package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// User represents the core user domain model.
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthTokens encapsulates issued JWT tokens.
type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// UserRepository defines the persistence contracts for User operations.
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
}

// UserService defines the business logic contracts for User & Auth operations.
type UserService interface {
	Register(ctx context.Context, name, email, password string) (*AuthTokens, error)
	Login(ctx context.Context, email, password string) (*AuthTokens, error)
	GetMe(ctx context.Context, userID uuid.UUID) (*User, error)
	RefreshToken(ctx context.Context, userID uuid.UUID) (*AuthTokens, error)
}
