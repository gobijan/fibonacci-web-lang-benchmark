defmodule OnPlug do
  use Application

  def start(_type, _args) do
    {:ok, _} = Plug.Adapters.Cowboy.http AppRouter, []
  end
end
