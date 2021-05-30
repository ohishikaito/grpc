class UsersController < ApplicationController
  def index
    req = Pb::GetUsersReq.new
    users = Stubs::UserStub::Stub.get_users(req)
    render json: users, status: :ok
  end

  def create
    req = Pb::CreateUserReq.new({
      last_name: user_params[:last_name],
      first_name: user_params[:first_name]
    })
    user = Stubs::UserStub::Stub.create_user(req)
    render json: user, status: :created
  end

  def show
    req = Pb::GetUserReq.new({id: params[:id].to_i})
    user = Stubs::UserStub::Stub.get_user(req)
    render json: user, status: :ok
  end

  def update
    req = Pb::UpdateUserReq.new({
      id: user_params[:id],
      last_name: user_params[:last_name],
      first_name: user_params[:first_name]
    })
    user = Stubs::UserStub::Stub.update_user(req)
    render json: user, status: :ok
  end

  def destroy
    req = Pb::DestroyUserReq.new({id: params[:id].to_i})
    user = Stubs::UserStub::Stub.destroy_user(req)
    head :no_content
  end

  private
    def user_params
      params.require(:user).permit(
        :id,
        :last_name,
        :first_name
      )
    end
end
