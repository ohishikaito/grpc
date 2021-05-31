package converter

import (
	"grpc/domain"
	pb "grpc/pb/user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUsers(users []*domain.User) (*pb.Users, error) {
	var pbUsers []*pb.User
	for _, user := range users {
		pbUser, _ := ConvertUser(user)
		pbUsers = append(pbUsers, pbUser)
	}
	return &pb.Users{Users: pbUsers}, nil
}

func ConvertUser(user *domain.User) (*pb.User, error) {
	pbUser := &pb.User{
		Id:        user.Id,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
	return pbUser, nil
}
