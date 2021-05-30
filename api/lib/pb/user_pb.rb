# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

require 'google/protobuf'

require 'google/protobuf/timestamp_pb'
require 'google/protobuf/empty_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("user.proto", :syntax => :proto3) do
    add_message "pb.GetUsersReq" do
    end
    add_message "pb.GetUserReq" do
      optional :id, :uint64, 1
    end
    add_message "pb.CreateUserReq" do
      optional :last_name, :string, 1
      optional :first_name, :string, 2
    end
    add_message "pb.UpdateUserReq" do
      optional :id, :uint64, 1
      optional :last_name, :string, 2
      optional :first_name, :string, 3
    end
    add_message "pb.DestroyUserReq" do
      optional :id, :uint64, 1
    end
    add_message "pb.User" do
      optional :id, :uint64, 1
      optional :last_name, :string, 2
      optional :first_name, :string, 3
      optional :created_at, :message, 4, "google.protobuf.Timestamp"
      optional :updated_at, :message, 5, "google.protobuf.Timestamp"
    end
    add_message "pb.Users" do
      repeated :users, :message, 1, "pb.User"
    end
  end
end

module Pb
  GetUsersReq = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("pb.GetUsersReq").msgclass
  GetUserReq = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("pb.GetUserReq").msgclass
  CreateUserReq = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("pb.CreateUserReq").msgclass
  UpdateUserReq = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("pb.UpdateUserReq").msgclass
  DestroyUserReq = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("pb.DestroyUserReq").msgclass
  User = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("pb.User").msgclass
  Users = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("pb.Users").msgclass
end
