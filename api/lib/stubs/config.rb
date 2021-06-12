module Stubs
  module Config
    name = ENV['GRPC_SERVICE_NAME'].presence || "localhost"
    # name = 'piyo.mynote.world' # 本番ALBのDNS名
    credentials = File.read('./credentials/ca.crt')
    CHANNEL_CREDS = GRPC::Core::ChannelCredentials.new(credentials)
    port = ENV['GRPC_SERVICE_PORT'].presence || "50051"
    GRPC_HOSTNAME = name + ":" + port

    def create_stub(stub)
      stub.new(GRPC_HOSTNAME, CHANNEL_CREDS)
      # stub.new(GRPC_HOSTNAME, :this_channel_is_insecure)

      # stub_opts = { channel_args: { GRPC::Core::Channel::SSL_TARGET => "go_app_1" } }
      # stub.new("go_app_1:50051", GRPC::Core::ChannelCredentials.new(), **stub_opts)

      # options = {
      #   timeout: 10
      # }

      # case ENV['RAILS_ENVIRONMENT']
      # when "production"
      #   stub.new(GRPC_HOSTNAME, channel_creds)
      # else
      #   stub.new(GRPC_HOSTNAME, :this_channel_is_insecure)
      # end
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