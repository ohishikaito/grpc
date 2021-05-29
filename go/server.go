package main

import (
	"context"
	"fmt"
	"grpc/proto/pinger"
	"grpc/proto/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}

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

func (s *server) GetUser(ctx context.Context, req *user.Empty) (*user.User, error) {
	user := &user.User{
		Name: "user",
	}
	return user, nil
}
