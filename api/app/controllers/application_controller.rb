require 'stubs'

class ApplicationController < ActionController::API
  def ping
    pong = Proto::Stubs::PingerStub.ping(Pinger::Empty.new)
    render json: { pong: pong.text }
  end
end
