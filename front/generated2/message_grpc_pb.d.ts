// GENERATED CODE -- DO NOT EDIT!

// package: Posts
// file: message.proto

import * as message_pb from "./message_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as grpc from "@grpc/grpc-js";

interface IPostsService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  getPosts: grpc.MethodDefinition<google_protobuf_empty_pb.Empty, message_pb.GetPostsResponse>;
  addPost: grpc.MethodDefinition<message_pb.AddPostRequest, message_pb.AddPostResponse>;
}

export const PostsService: IPostsService;

export interface IPostsServer extends grpc.UntypedServiceImplementation {
  getPosts: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, message_pb.GetPostsResponse>;
  addPost: grpc.handleUnaryCall<message_pb.AddPostRequest, message_pb.AddPostResponse>;
}

export class PostsClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  getPosts(argument: google_protobuf_empty_pb.Empty, callback: grpc.requestCallback<message_pb.GetPostsResponse>): grpc.ClientUnaryCall;
  getPosts(argument: google_protobuf_empty_pb.Empty, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<message_pb.GetPostsResponse>): grpc.ClientUnaryCall;
  getPosts(argument: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<message_pb.GetPostsResponse>): grpc.ClientUnaryCall;
  addPost(argument: message_pb.AddPostRequest, callback: grpc.requestCallback<message_pb.AddPostResponse>): grpc.ClientUnaryCall;
  addPost(argument: message_pb.AddPostRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<message_pb.AddPostResponse>): grpc.ClientUnaryCall;
  addPost(argument: message_pb.AddPostRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<message_pb.AddPostResponse>): grpc.ClientUnaryCall;
}
