require 'config'

module Stubs
  module UserStub
    Stub = Pb::User::UserService::Stub.new(Stubs::Config::GRPC_HOSTNAME, :this_channel_is_insecure)
  end
end