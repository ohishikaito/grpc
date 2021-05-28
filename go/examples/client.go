package main

import (
	"context"
	"fmt"
	"grpc/proto/pinger"
	"os"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5300", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "grpc.Dial: %v\n", err)
		return
	}
	defer conn.Close()
	client := pinger.NewPingerServiceClient(conn)
	req := &pinger.Empty{}
	pong, err := client.Ping(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping: %v\n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "Pong: %s\n", pong.Text)
}
