module Stubs
  module Config
    # 環境変数を定義
    name = ENV['GRPC_SERVICE_NAME']
    port = ENV['GRPC_SERVICE_PORT']
    # name = "localhost"
    # port = "5300"
    GRPC_HOSTNAME = name + ":" + port

    # stubを生成するfacotry pattern
    def create_stub(stub)
      stub.new(GRPC_HOSTNAME, :this_channel_is_insecure)
    end
  end
end