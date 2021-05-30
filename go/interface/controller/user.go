package controller

import (
	"context"
	"grpc/interface/converter"
	pb "grpc/protos/user"
	"grpc/usecase"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) *userController {
	return &userController{
		userUsecase: uu,
	}
}

func (c *userController) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.Users, error) {
	users, err := c.userUsecase.GetUsers()
	if err != nil {
		return nil, err
	}
	return converter.ConvertUsers(users)

}

func (c *userController) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.User, error) {
	user, err := c.userUsecase.GetUserById(req.Id)
	if err != nil {
		return nil, err
	}
	return converter.ConvertUser(user)
}
