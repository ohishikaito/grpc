package main

import (
	"fmt"
	"grpc/infrastructure"
	"grpc/infrastructure/repository"
	"grpc/interface/controller"
	"grpc/protos/user"
	"grpc/usecase"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	db := infrastructure.NewGormConnect()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, userController)

	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_SERVICE_PORT")) // [::]:5300
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	fmt.Printf("grpcServer.Serve")

	grpcServer.Serve(listener)
}
