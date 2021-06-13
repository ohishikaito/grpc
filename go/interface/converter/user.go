package converter

import (
	"grpc/domain"
	pb "grpc/pb/user"
)

func ConvertUsers(users []*domain.User) (*pb.GetUsersResponse, error) {
	var pbUsers []*pb.User
	for _, user := range users {
		pbUser, _ := ConvertUser(user)
		pbUsers = append(pbUsers, pbUser)
	}
	return &pb.GetUsersResponse{Users: pbUsers}, nil
}

func ConvertUser(user *domain.User) (*pb.User, error) {
	pbUser := &pb.User{
		Id:        user.Id,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		CreatedAt: FromUTCtoJSTinProto(user.CreatedAt),
		UpdatedAt: FromUTCtoJSTinProto(user.UpdatedAt),
	}
	return pbUser, nil
}
