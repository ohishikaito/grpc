package infrastructure

import (
	"log"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func NewGrpcServer() *grpc.Server {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	switch os.Getenv("environment") {
	case "production":
		// NOTE: client側はcrtファイルいらないけど、server側はTLS化する必要あり
		creds, err := credentials.NewServerTLSFromFile(
			"credentials/ca.crt",
			"credentials/server.key",
		)
		if err != nil {
			log.Fatalf("failed to load certificate: %v", err)
		}
		return grpc.NewServer(
			grpc.Creds(creds),
			grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
				grpc_zap.StreamServerInterceptor(zapLogger),
				grpc_recovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(zapLogger),
				grpc_recovery.UnaryServerInterceptor(),
			)),
		)
	default:
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
		reflection.Register(grpcServer)
		return grpcServer
	}
}
