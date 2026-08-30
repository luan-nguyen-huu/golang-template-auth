package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/luan-nguyen-huu/Adam/internal/domain"
)

// User is the GORM database schema entity for persistence.
type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// ToDomain converts the persistence User entity to a pure domain.User model.
func (u *User) ToDomain() *domain.User {
	if u == nil {
		return nil
	}
	return &domain.User{
		ID:        u.ID,
		Email:     u.Email,
		Password:  u.PasswordHash,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// FromDomainUser creates a persistence User entity from a pure domain.User model.
func FromDomainUser(u *domain.User) *User {
	if u == nil {
		return nil
	}
	return &User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.Password,
		Name:         u.Name,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}