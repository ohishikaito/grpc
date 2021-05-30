package converter

import (
	"grpc/domain"
	pb "grpc/protos/user"
)

func ConvertUsers(users []*domain.User) (*pb.Users, error) {
	var pbUsers []*pb.User
	for _, user := range users {
		pbUser, err := ConvertUser(user)
		if err != nil {
			return nil, err
		}
		pbUsers = append(pbUsers, pbUser)
	}
	return &pb.Users{Users: pbUsers}, nil
}

func ConvertUser(user *domain.User) (*pb.User, error) {
	pbUser := &pb.User{
		Id:        user.Id,
		LastName:  user.LastName,
		FirstName: user.FirstName,
	}
	return pbUser, nil
}
