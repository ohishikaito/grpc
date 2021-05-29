class UsersController < ApplicationController
  def index
    req = UserProto::Empty.new
    users = Protos::Stubs::UserStub.get_users(req)
    render json: users, status: :ok
  end

  def show
    req = UserProto::GetUserReq.new({id: 11})
    user = Protos::Stubs::UserStub.get_user(req)
    render json: user, status: :ok
  end
end
