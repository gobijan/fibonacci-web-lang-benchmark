-module(fib_app).

-behaviour(application).

-export([start/0]).
-export([start/2, stop/1]).

start() ->
    application:ensure_all_started(fib).

start(_StartType, _StartArgs) ->
    Dispatch = cowboy_router:compile([
        {'_', [
            {"/:number", fib_index_handler, []}]}]),
    {ok, _} = cowboy:start_http(http, 100, [{port, 8088}],
                    [{env, [{dispatch, Dispatch}]}]),
    fib_sup:start_link().

stop(_State) ->
    ok.
