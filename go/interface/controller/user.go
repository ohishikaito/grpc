package controller

import (
	"context"
	"grpc/protos/user"
	"grpc/usecase"
)

type UserController struct {
	userUsercase usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) *UserController {
	return &UserController{
		userUsercase: uu,
	}
}

func (c *UserController) GetUser(ctx context.Context, req *user.Empty) (*user.User, error) {
	// 後で作る
	// user, err := c.userUsercase.GetUser()
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}
