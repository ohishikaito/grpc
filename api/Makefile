attach:
	docker attach api_app_1

bash:
	docker exec -it api_app_1 bash

go-pb-gen:
	protoc -I ./protofiles ${s}.proto --go_out=plugins=grpc:../go/pb

ruby-pb-ruby:
	docker-compose run app \
	bundle exec grpc_tools_ruby_protoc -I ./protofiles \
	--ruby_out=lib/pb --grpc_out=lib/pb ./protofiles/*.proto

pb-gen:
	protoc -I ./protofiles ${s}.proto --go_out=plugins=grpc:../go/pb && \
	docker-compose run --rm app \
	bundle exec grpc_tools_ruby_protoc -I ./protofiles \
		--ruby_out=lib/pb --grpc_out=lib/pb ./protofiles/${s}.proto && \
	docker-compose restart app

push:
	docker build -t grpc-rails -f ./docker/api/Dockerfile . && \
	docker tag grpc-rails:latest 504300096460.dkr.ecr.ap-northeast-1.amazonaws.com/grpc-rails:latest && \
	docker push 504300096460.dkr.ecr.ap-northeast-1.amazonaws.com/grpc-rails:latest