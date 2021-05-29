package main

import (
	"context"
	"fmt"
	"grpc/protos/pinger"
	"grpc/protos/user"
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

	pinger.RegisterPingerServiceServer(grpcServer, &server{})
	user.RegisterUserServiceServer(grpcServer, &server{})

	fmt.Printf("pinger && user!!!!!!!")
	grpcServer.Serve(listener)
}

func (s *server) Ping(ctx context.Context, req *pinger.Empty) (*pinger.Pong, error) {
	pong := &pinger.Pong{
		Text: "pong",
	}
	return pong, nil
}

func (s *server) GetUsers(ctx context.Context, req *user.Empty) (*user.User, error) {
	user := &user.User{
		Name: "kaito",
	}
	return user, nil
}

func (s *server) GetUser(ctx context.Context, req *user.GetUserReq) (*user.User, error) {
	fmt.Println(req, "req")
	user := &user.User{
		Name: "getUser",
	}
	return user, nil
}
