Rails.application.routes.draw do
  resources :users, only: %i[index create show update destory]
end
