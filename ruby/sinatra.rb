require 'sinatra'

def fib(n)	
	if n == 0
		return 0
	elsif n == 1
		return 1
	else
		return fib(n-1) + fib(n-2)
	end
end

get '/:number' do	
	"#{fib(params[:number].to_i)}"
end