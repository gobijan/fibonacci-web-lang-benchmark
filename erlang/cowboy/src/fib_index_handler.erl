-module(fib_index_handler).
-behaviour(cowboy_http_handler).

-export([init/3, handle/2, terminate/3]).

init(_Transport, Req, _Opts) ->
    {ok, Req, []}.

handle(Req, State) ->
    {BinNumber, Req1} = cowboy_req:binding(number, Req),
    Fib = fib(binary_to_integer(BinNumber)),
    BinFib = integer_to_binary(Fib),
    Content = [<<"Erlang + Cowboy<hr>fib(">>, BinNumber, <<"): ">>, BinFib],
    {ok, Resp} = cowboy_req:reply(200,
                            [{<<"content-type">>, <<"text/html">>}],
                            Content,
                            Req1),
    {ok, Resp, State}.

terminate(_Reason, _Req, _State) ->
    ok.

fib(0) -> 0;
fib(1) -> 1;
fib(N) when N > 0 ->
    fib(N - 1) + fib(N - 2).
