<!-- root dirで
`protoc -I ./proto pinger.proto --go_out=plugins=grpc:./pinger/lib`
ってやってたけど、↓のerror
 
```
protoc-gen-go: unable to determine Go import path for "pinger.proto"

Please specify either:
	• a "go_package" option in the .proto source file, or
	• a "M" argument on the command line.

See https://developers.google.com/protocol-buffers/docs/reference/go-generated#package for more information.

--go_out: protoc-gen-go: Plugin failed with status code 1.
```

./proto/pinger.protoに、
`option go_package = "./pinger";`
って書いて、
`protoc -I ./proto pinger.proto --go_out=.`
ってコマンド叩いたら、ファイル生成された。くそ。なんなんまじで？


環境変数には↓　コマンドの意味は分からんから調べる
`export GO_PATH=~/go`
`export PATH=$PATH:/$GO_PATH/bin`
https://stackoverflow.com/questions/28099004/unable-to-build-protobuf-to-go-endpoint

`echo $PATH`すると↓
/Users/ohishikaido/google-cloud-sdk/bin:/Users/ohishikaido/.rbenv/shims:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/usr/local/share/dotnet:~/.dotnet/tools:/Library/Frameworks/Mono.framework/Versions/Current/Commands://Users/ohishikaido/go/bin

結果的に
`./proto/pinger.proto`に
`option go_package = "./pinger";`
って書いて、
root dirで
`protoc -I ./proto pinger.proto --go_out=plugins=grpc:./go`
ってコマンド叩いたら、`./go/pinger`にgRPCのファイル生成される

`option go_package = "./";`のヒント
https://github.com/evilsocket/opensnitch/issues/373

current dirが`server $`で、
`bundle exec grpc_tools_ruby_protoc -I ../proto --ruby_out=lib --grpc_out=lib ../proto/pinger.proto`
って叩いたら、protoファイルが生成される！
`bundle exec grpc_tools_ruby_protoc -I ../proto --ruby_out=lib/pb --grpc_out=lib/pb ../proto/pinger.proto`

またはrakeタスクを使って、
`rake pb:protoc`でもok➡️これはsmartHRのパクリ -->
