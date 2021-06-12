push:
	docker build -t grpc-go -f ./docker/Dockerfile  .
	docker tag grpc-go:latest 504300096460.dkr.ecr.ap-northeast-1.amazonaws.com/grpc-go:latest
	docker push 504300096460.dkr.ecr.ap-northeast-1.amazonaws.com/grpc-go:latest