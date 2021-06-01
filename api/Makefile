attach:
	docker attach api_app_1

bash:
	docker exec -it api_app_1 /bin/bash

restart:
	docker-compose down && \
	docker-compose up

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