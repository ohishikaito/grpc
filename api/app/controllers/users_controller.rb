class UsersController < ApplicationController
  def index
    user = Protos::Stubs::UserStub.get_user(UserProto::Empty.new)
    render json: { user: user.name }
  end
end
