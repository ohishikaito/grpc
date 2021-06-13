package main

import (
	"grpc/infrastructure"
	"grpc/infrastructure/repository"
	"grpc/interface/controller"
	"grpc/pb/user"
	"grpc/usecase"
	"log"
	"net"
	"os"
)

func main() {
	db := infrastructure.NewGormConnect()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	grpcServer := infrastructure.NewGrpcServer()
	user.RegisterUserServiceServer(grpcServer, userController)

	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_SERVICE_PORT")) // [::]:50051
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	log.Print("grpcServer serve")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("serve err", err)
	}
}
