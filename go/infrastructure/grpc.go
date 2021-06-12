package infrastructure

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewGrpcServer() *grpc.Server {
	switch os.Getenv("environment") {
	case "production":
		// go側で認証はしなくても良い説？
		creds, err := credentials.NewServerTLSFromFile(
			"credentials/ca.crt",
			"credentials/server.key",
		)
		if err != nil {
			log.Fatalf("failed to load certificate: %v", err)
		}
		return grpc.NewServer(grpc.Creds(creds))
	default:
		return grpc.NewServer()
	}
}
