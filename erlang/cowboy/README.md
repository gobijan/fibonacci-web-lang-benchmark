
# Erlang Cowboy Version of the Fibonacci Web Lang Benchmark

Author: Yüce Tekol <yucetekol@gmail.com>

## Info

This sample uses version 1.0.1 of [Cowboy](https://github.com/ninenines/cowboy.git).

The router is at `src/fib_app.erl`. Cowboy can constrain the arguments passed to the handler (*number* in this case), but it causes the number of requests per second to drop about 1000 on my machine, so I didn't use it.

The handler is `src/fib_index_handler.erl`.

The proper way to run an OTP based Erlang application is creating a release first, but that's unnecessary for just trying out this simple code.

## Instructions

Compile with:

    make compile

and then run with:

    make run

The app runs at `localhost:8088` by default.

You can exit the Erlang shell by pressing Ctrl+C twice.
