REBAR = ./rebar

.PHONY: all deps compile clean

all: deps compile

run:
	erl  -pa ./deps/cowboy/ebin ./deps/cowlib/ebin ./deps/ranch/ebin ./ebin -s fib_app

deps:
	$(REBAR) get-deps

compile:
	$(REBAR) compile

clean:
	$(REBAR) clean
	rm -f erl_crash.dump
