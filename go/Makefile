push:
	aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 504300096460.dkr.ecr.ap-northeast-1.amazonaws.com
	docker build -t grpc-go -f ./docker/Dockerfile  .
	docker tag grpc-go:latest 504300096460.dkr.ecr.ap-northeast-1.amazonaws.com/grpc-go:latest
	docker push 504300096460.dkr.ecr.ap-northeast-1.amazonaws.com/grpc-go:latest

build-run:
	docker build -t grpc-go -f ./docker/Dockerfile  .
	docker run grpc-go

curl:
	grpcurl -plaintext localhost:50051  pb.UserService/GetUsers

curl2:
	grpcurl -plaintext -d '{"id": 10 }' localhost:50051  pb.UserService/GetUser