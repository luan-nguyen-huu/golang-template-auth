package repositories

import (
	"github.com/luan-nguyen-huu/Adam/internal/entities"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entities.UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(userID uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}