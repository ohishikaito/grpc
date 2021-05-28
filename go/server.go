package main

import (
	"context"
	"fmt"
	"grpc/proto/pinger"
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
	fmt.Printf("pinger!!!!!!!")
	grpcServer.Serve(listener)
}

func (s *server) Ping(ctx context.Context, req *pinger.Empty) (*pinger.Pong, error) {
	pong := &pinger.Pong{
		Text: "pong",
	}
	return pong, nil
}
