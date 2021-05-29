class UsersController < ApplicationController
  # userを配列でとれるようにしたい
  def index
    users = Protos::Stubs::UserStub.get_users(UserProto::Empty.new)
    # render json: { users: users }
    render json: { users: users.name }
  end

  def show
    req = UserProto::GetUserReq.new({id: 11})
    user = Protos::Stubs::UserStub.get_user(req)
    render json: { user: user.name }
  end
end
