class RootController < ApplicationController
  def index
    render json: {
      hoge: "fuga",
      foo: 1000000000
    }
  end
end
