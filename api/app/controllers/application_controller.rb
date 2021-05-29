require 'pinger_services_pb.rb'
require 'pinger_pb.rb'

class ApplicationController < ActionController::API
  def ping
    name = ENV['GRPC_SERVICE_NAME']
    port = ENV['GRPC_SERVICE_PORT']
    pinger_stub = Pinger::PingerService::Stub.new(name + ":" + port, :this_channel_is_insecure)
    pong = pinger_stub.ping(Pinger::Empty.new)
    render json: {pong: pong.text}
  end
end
