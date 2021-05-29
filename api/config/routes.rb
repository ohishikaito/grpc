Rails.application.routes.draw do
  get 'ping', to: 'application#ping'
  resources :users, only: %i[index create show update destory]
end
