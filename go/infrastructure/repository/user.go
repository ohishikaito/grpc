package repository

import (
	"grpc/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUsers() ([]*domain.User, error)
	GetUserById(id uint64) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUsers() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserById(id uint64) (*domain.User, error) {
	user := &domain.User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
