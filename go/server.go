package main

import (
	"context"
	"fmt"
	pb "grpc/protos/user"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_SERVICE_PORT")) // [::]:5300
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	// conn := infrastructure.NewGormConnect()
	// userRepository := repository.NewUserRepository(conn)
	// userUsecase := usecase.NewUserUsecase(userRepository)
	// userController := controller.NewUserController(userUsecase)
	// user.RegisterUserServiceServer(grpcServer, userController)

	grpcServer := grpc.NewServer()

	// pinger.RegisterPingerServiceServer(grpcServer, &server{})
	pb.RegisterUserServiceServer(grpcServer, &server{})

	fmt.Printf("pinger && user!!!!!!!")
	grpcServer.Serve(listener)
}

// 今はpbで指定してるけど、controllerに移行したら問題なくなる
func (s *server) GetUsers(ctx context.Context, req *pb.Empty) (*pb.Users, error) {
	var pbUsers []*pb.User
	pbUser := &pb.User{
		LastName:  "lastName",
		FirstName: "firstName",
		Email:     "email",
	}
	pbUsers = append(pbUsers, pbUser)
	pbUsers = append(pbUsers, pbUser)
	pbUsers = append(pbUsers, pbUser)
	fmt.Println(pbUsers)
	return &pb.Users{
		Users: pbUsers,
	}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.User, error) {
	fmt.Println(req, "req")
	user := &pb.User{
		LastName:  "lastName",
		FirstName: "firstName",
		Email:     "email",
	}
	return user, nil
}
