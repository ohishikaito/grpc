package repository

import (
	"grpc/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUser() (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUser() (*domain.User, error) {
	var user *domain.User
	if err := r.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
