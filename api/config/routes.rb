Rails.application.routes.draw do
  resources :posts
  resources :users, only: %i[index create show update destroy]
end
