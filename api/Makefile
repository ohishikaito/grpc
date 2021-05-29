bash:
	docker exec -it api_app_1 /bin/bash

protoc:
	docker-compose run app \
	bundle exec grpc_tools_ruby_protoc -I ./protofiles \
	--ruby_out=lib/proto --grpc_out=lib/proto ./protofiles/*.proto