package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID               uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email            string         `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash     string         `gorm:"not null" json:"-"`
	Name             string         `json:"name"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	IsDeleted        bool           `gorm:"default:false" json:"is_deleted"`
	IsVerified       bool           `gorm:"default:false" json:"is_verified"`
	Active           int            `gorm:"default:0" json:"active"`
	DurationRegister time.Time      `gorm:"autoUpdateTime;not null" json:"duration_register"`
}


type UserRepositoryInterface interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByID(userID uuid.UUID) (*User, error)
}

type UserServiceInterface interface {
	RegisterUser(username string, password string, email string) (string, string, error)
	LoginUser(email string, password string) (string, string, error)
	GetMe(userID uuid.UUID) (*User, error)
	RefreshToken(userID uuid.UUID) (string, string, error)
}