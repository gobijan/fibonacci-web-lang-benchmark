defmodule AppRouter do
  use Plug.Router

  plug :match
  plug :dispatch
  

  get "/:n" do
    n = String.to_integer(n)
    conn |> send_resp 200, to_string(Fibonacci.get(n))
  end
end