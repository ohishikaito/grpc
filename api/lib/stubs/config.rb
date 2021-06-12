module Stubs
  module Config
    name = ENV['GRPC_SERVICE_NAME'].presence || "localhost"
    name = 'piyo.mynote.world' # 本番ALBのDNS名

    # NOTE: 空の引数を渡すと何故か動く。理由は分からん。
    port = ENV['GRPC_SERVICE_PORT'].presence || "50051"
    GRPC_HOSTNAME = name + ":" + port

    def create_stub(stub)
      # stub_opts = { channel_args: { GRPC::Core::Channel::SSL_TARGET => "go_app_1" } }
      # stub.new("go_app_1:50051", GRPC::Core::ChannelCredentials.new(), **stub_opts)

      case ENV['RAILS_ENV']
      when "production"
        stub.new(GRPC_HOSTNAME, GRPC::Core::ChannelCredentials.new())
      else
        # stub.new(GRPC_HOSTNAME, :this_channel_is_insecure)
        stub.new(GRPC_HOSTNAME, GRPC::Core::ChannelCredentials.new())
      end
    end
  end
end

# GoのDebugger
require 'grpc'

module GrpcLogger
  def logger
    LOGGER
  end
  LOGGER = Logger.new(STDOUT)
  LOGGER.level = :debug

end

module GRPC
  extend GrpcLogger
end