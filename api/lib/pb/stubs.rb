require 'user_pb.rb'
require 'user_services_pb.rb'

module Pb
  module Stubs
    # 環境変数を定義
    name = ENV['GRPC_SERVICE_NAME']
    port = ENV['GRPC_SERVICE_PORT']
    hostname = name + ":" + port

    UserStub = Pb::User::UserService::Stub.new(hostname, :this_channel_is_insecure)
  end
end

# ファイル名をUserStubにして、定数名をStubにする