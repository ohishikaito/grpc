package main

import (
	"context"
	"fmt"
	"grpc/pb/user"

	"google.golang.org/grpc"
)

// NOTE: gRPCのリクエストを試すようのclient
func main() {
	request()
}

func request() {
	conn := newGrpcClientConn()
	client := user.NewUserServiceClient(conn)

	req := &user.GetUserReq{Id: 1}
	result, err := client.GetUser(context.Background(), req)
	if err != nil {
		panic(err)
	}
	fmt.Println("--------- response data ---------")
	fmt.Println(result)
	fmt.Println("---------------------------------")
}

func newGrpcClientConn() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}
