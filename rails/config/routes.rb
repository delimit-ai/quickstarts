Rails.application.routes.draw do
  get "health", to: "api/tasks#health"
  resources :tasks, only: [:index, :create, :show, :destroy], controller: "api/tasks"
end
