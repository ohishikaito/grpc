package converter

import (
	"grpc/domain"
	pb "grpc/pb/user"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
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
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		// convertしなくてもrailsがUTCをJTCに変えてくれるんでよさげ〜
		// CreatedAt:      FromUTCtoJSTinProto(user.CreatedAt),
		// UpdatedAt:      FromUTCtoJSTinProto(user.UpdatedAt),
		Liked:          false,
		BazirisukuTime: timestamppb.New(time.Now()),
		OrderDate:      timestamppb.New(time.Now()),
	}
	return pbUser, nil
}
