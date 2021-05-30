# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: user.proto for package 'pb'

require 'grpc'
require 'user_pb'

module Pb
  module UserService
    class Service

      include ::GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'pb.UserService'

      rpc :GetUsers, ::Pb::GetUsersReq, ::Pb::Users
      rpc :GetUser, ::Pb::GetUserReq, ::Pb::User
      rpc :CreateUser, ::Pb::CreateUserReq, ::Pb::User
      rpc :UpdateUser, ::Pb::UpdateUserReq, ::Pb::User
      rpc :DestroyUser, ::Pb::DestroyUserReq, ::Google::Protobuf::Empty
    end

    Stub = Service.rpc_stub_class
  end
end
