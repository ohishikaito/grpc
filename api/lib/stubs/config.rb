module Stubs
  module Config
    name = ENV['GRPC_SERVICE_NAME'].presence || "localhost"
    port = ENV['GRPC_SERVICE_PORT'].presence || "50051"
    GRPC_HOSTNAME = name + ":" + port

    def create_stub(stub)
      case ENV['RAILS_ENV']
      when "production"
        stub.new(GRPC_HOSTNAME, GRPC::Core::ChannelCredentials.new())
      else
        stub.new("piyo.mynote.world:50051", GRPC::Core::ChannelCredentials.new())
        stub.new(GRPC_HOSTNAME, :this_channel_is_insecure)
      end
    end
  end
end

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