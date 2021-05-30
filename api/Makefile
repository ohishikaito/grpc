attach:
	docker attach api_app_1

bash:
	docker exec -it api_app_1 /bin/bash

protoc:
	docker-compose run app \
	bundle exec grpc_tools_ruby_protoc -I ./protofiles \
	--ruby_out=lib/pb --grpc_out=lib/pb ./protofiles/*.proto