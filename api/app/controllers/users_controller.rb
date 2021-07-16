class UsersController < ApplicationController
  def index
    req = Google::Protobuf::Empty.new
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
    pb_instance = Stubs::UserStub::Stub.get_user(req)
    # 手動でモデル作るっていう手段もできるけど、くそめんどいからやりたくない
    # hoge = converter(pb_instance)

    god_user = initialize_of_model(User, pb_instance) do |instance_hash|
      instance_hash[:order_date] = convert_epoc_time_to_time_class(pb_instance.order_date.seconds)
    end
    render json: god_user, status: :ok
  end

  def initialize_of_model(model, pb_instance, &block)
    instance_hash = pb_instance.to_h
    instance_hash[:created_at] = convert_epoc_time_to_time_class(pb_instance.created_at.seconds)
    instance_hash[:updated_at] = convert_epoc_time_to_time_class(pb_instance.updated_at.seconds)
    yield(instance_hash) if block_given?
    model.new(instance_hash)
  end

  def convert_epoc_time_to_time_class(epoc_time)
    Time.at(epoc_time).in_time_zone('Tokyo')
  end

  def converter(pb_user)
    User.new({
      id: pb_user.id,
      last_name: pb_user.last_name,
      first_name: pb_user.first_name,
      epoc_created_at: pb_user.created_at.seconds,
      # updated_at: pb_user.updated_at,
      liked: true
    })
  end

  def update
    req = Pb::UpdateUserReq.new({
      # NOTE: user_paramsからidを取得すると、別のユーザーを更新できる&idがないと全件更新される
      # id: user_params[:id],
      id: params[:id].to_i,
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
