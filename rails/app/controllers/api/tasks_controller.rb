module Api
  class TasksController < ActionController::API
    @@tasks = {}
    @@next_id = 1

    def health
      render json: { status: "ok" }
    end

    def index
      render json: @@tasks.values
    end

    def create
      id = @@next_id
      @@next_id += 1
      task = {
        id: id,
        title: params[:title],
        description: params.fetch(:description, ""),
        done: false
      }
      @@tasks[id] = task
      render json: task, status: :created
    end

    def show
      task = @@tasks[params[:id].to_i]
      if task
        render json: task
      else
        head :not_found
      end
    end

    def destroy
      if @@tasks.delete(params[:id].to_i)
        head :no_content
      else
        head :not_found
      end
    end
  end
end
