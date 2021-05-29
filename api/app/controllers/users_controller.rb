class UsersController < ApplicationController
  # userを配列でとれるようにしたい
  def index
    users = Protos::Stubs::UserStub.get_users(UserProto::Empty.new)
    # render json: { users: users }
    p users.users[0]
    # binding.pry
    # render json: { users: users.users }
    render json: { users: users.users[0].to_json }
    # render json: { users: users.name }
  end

  def show
    req = UserProto::GetUserReq.new({id: 11})
    p req
    user = Protos::Stubs::UserStub.get_user(req)
    p user
    render json: { user: user.email }
  end
end
