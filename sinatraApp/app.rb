require "sinatra"
require "json"

get '/' do
  content_type :json
  { hoge: "fuga", foo: 1000000000 }.to_json
end
