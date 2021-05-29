require 'user_proto_pb.rb'
require 'user_proto_services_pb.rb'

module Protos
  module Stubs
    # 環境変数を定義
    name = ENV['GRPC_SERVICE_NAME']
    port = ENV['GRPC_SERVICE_PORT']
    hostname = name + ":" + port

    UserStub = UserProto::UserService::Stub.new(hostname, :this_channel_is_insecure)
  end
end