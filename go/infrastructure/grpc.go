package infrastructure

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewGrpcServer() *grpc.Server {
	// go側で認証はしなくても良い説？
	creds, err := credentials.NewServerTLSFromFile(
		"credentials/ca.crt",
		"credentials/server.key",
	)
	if err != nil {
		fmt.Println("creds err", err)
		log.Fatalf("failed to load certificate: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
	)
	fmt.Println("grpcServer", grpcServer)
	return grpcServer

	// return grpc.NewServer()
}
