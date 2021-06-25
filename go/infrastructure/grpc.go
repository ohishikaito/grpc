package infrastructure

import (
	"fmt"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer() *grpc.Server {
	fmt.Println("更新されてるかーい？w")

	var zapLogger *zap.Logger
	switch os.Getenv("environment") {
	case "production":
		zapLogger, _ = zap.NewProduction()
	default:
		zapLogger, _ = zap.NewDevelopment()
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(zapLogger),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	// NOTE: grpcurl用 development環境でのみ実行する
	if os.Getenv("environment") == "development" {
		reflection.Register(grpcServer)
	}

	return grpcServer
}
