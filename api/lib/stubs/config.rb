module Stubs
  module Config
    # 環境変数を定義
    # dockerの時はENVから取得する。localの時はlocalhostから。
    name = ENV['GRPC_SERVICE_NAME'].presence || "localhost"
    port = ENV['GRPC_SERVICE_PORT'].presence || "50051"
    GRPC_HOSTNAME = name + ":" + port

    # stubを生成するfacotry pattern
    def create_stub(stub)
      stub.new(GRPC_HOSTNAME, :this_channel_is_insecure)
    end
  end
end