require 'pinger_pb.rb'
require 'pinger_services_pb.rb'
require 'user_pb.rb'
require 'user_services_pb.rb'

module Proto
  module Stubs
    # 環境変数を定義
    name = ENV['GRPC_SERVICE_NAME']
    port = ENV['GRPC_SERVICE_PORT']
    hostname = name + ":" + port

    # pingerだけpackage名がpingerになってる
    PingerStub = Pinger::PingerService::Stub.new(hostname, :this_channel_is_insecure)

    UserStub = Proto::UserService::Stub.new(hostname, :this_channel_is_insecure)
  end
end