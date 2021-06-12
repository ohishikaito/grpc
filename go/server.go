package main

import (
	"fmt"
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
	fmt.Println(db, "db")

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	grpcServer := infrastructure.NewGrpcServer()
	fmt.Println(grpcServer, "grpcServer")
	user.RegisterUserServiceServer(grpcServer, userController)
	fmt.Println("RegisterUserServiceServerしました")

	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_SERVICE_PORT")) // [::]:50051
	fmt.Println("Listenでerror", err)
	if err != nil {
		fmt.Println("owaowari")
		log.Fatalf("failed to listen: %v\n", err)
		return
	}
	fmt.Println("serveします！")
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("serve err", err)
	}
	fmt.Println("serveしました！")
}
