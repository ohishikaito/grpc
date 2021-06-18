codegen:
	yarn grpc_tools_node_protoc -I ./proto \
	--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
	--js_out=import_style=commonjs,binary:./generated \
	--grpc_out=grpc_js:./generated \
	--ts_out=service=grpc-node,mode=grpc-js:./generated ./proto/*.proto

# ESmodules形式にしてもかわんなそう笑
gen:
	yarn grpc_tools_node_protoc -I ./proto \
	--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
	--js_out=import_style=esmodules,binary:./generated2 \
	--grpc_out=grpc_js:./generated2 \
	--ts_out=service=grpc-node,mode=grpc-js:./generated2 ./proto/*.proto

# "scripts": {
#   "codegen": "grpc_tools_node_protoc -I ./proto
# 	--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts
# commonJS形式は怪しいかも
# 	--js_out=import_style=commonjs,binary:./generated
# 	--grpc_out=grpc_js:./generated
# 	--ts_out=service=grpc-node,mode=grpc-js:./generated ./proto/*.proto"
# },
#  grpc_tools_node_protoc -I <.protoファイルのディレクトリ> \
#    --plugin=<プラグイン名>=<プラグインのパス> --js_out=<オプション>:<クライアントコード生成先パス> \
#    --grpc_out=<オプション>:<サーバーコード生成先パス>  \
#    --ts_out=<オプション>:<TypeScript型定義生成先パス> \
#    <生成元のProtocol Buffersのパス>