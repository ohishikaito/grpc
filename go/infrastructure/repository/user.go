package repository

import (
	"grpc/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetUsers() ([]*domain.User, error)
	GetUserById(id uint64) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUserById(id uint64) error
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
		// 意味ある？
		// if gorm.IsRecordNotFoundError(err) {
		// 	err = status.Errorf(codes.NotFound, "user with id='%d' is not found", id)
		// }
		return nil, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	if err := r.db.Model(user).Update(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) DeleteUserById(id uint64) error {
	query := "DELETE FROM users WHERE id = ?"
	if err := r.db.Exec(query, id).Error; err != nil {
		return err
	}
	return nil
}
