Elixir + Plug
======

```
$ mix do deps.get, compile
$ MIX_ENV=prod iex -S mix
$ wrk -c 64 -d 30s http://localhost:4000/10 
```