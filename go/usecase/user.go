package usecase

import (
	"grpc/domain"
	"grpc/infrastructure/repository"
)

type UserUsecase interface {
	GetUser() (*domain.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (u *userUsecase) GetUser() (*domain.User, error) {
	return u.userRepository.GetUser()
}
