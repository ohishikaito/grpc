package usecase

import (
	"grpc/domain"
	"grpc/infrastructure/repository"
)

type UserUsecase interface {
	GetUsers() ([]*domain.User, error)
	GetUserById(id uint64) (*domain.User, error)
	CreateUser(user *domain.User) (*domain.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (u *userUsecase) GetUsers() ([]*domain.User, error) {
	return u.userRepository.GetUsers()
}

func (u *userUsecase) GetUserById(id uint64) (*domain.User, error) {
	return u.userRepository.GetUserById(id)
}

func (u *userUsecase) CreateUser(user *domain.User) (*domain.User, error) {
	return u.userRepository.CreateUser(user)
}
