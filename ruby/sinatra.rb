require 'sinatra'
set :logging, false

def fib(n)
  if n.zero?
    return 0
  elsif n == 1
    return 1
  else
    return fib(n - 1) + fib(n - 2)
  end
end

get '/:number' do
  fib(params[:number].to_i).to_s
end
