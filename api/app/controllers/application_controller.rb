require 'pinger_services_pb.rb'
require 'pinger_pb.rb'

class ApplicationController < ActionController::API
  def ping
    pinger_stub = Pinger::PingerService::Stub.new('172.24.0.2:5300', :this_channel_is_insecure)
    # pinger_stub = Pinger::PingerService::Stub.new('localhost:5300', :this_channel_is_insecure)
    # pinger_stub = Pinger::PingerService::Stub.new('0.0.0.0:5300', :this_channel_is_insecure)
    # pinger_stub = Pinger::PingerService::Stub.new('127.0.0.1:5300', :this_channel_is_insecure)
    # pinger_stub = Pinger::PingerService::Stub.new('http://localhost:5300', :this_channel_is_insecure) #=> error
    pong = pinger_stub.ping(Pinger::Empty.new)
    render json: {pong: pong.text}
  end
end
