require 'stubs'

class ApplicationController < ActionController::API
  def ping
    # pong = Proto::Stubs::PingerStub.ping(Pinger::Empty.new)
    # render json: { pong: pong.text }

    # 仮でpingに作る
    user = Proto::Stubs::UserStub.get_user(Proto::Empty.new)
    render json: { user: user.name }
    # render json: { user: user }
  end
end
