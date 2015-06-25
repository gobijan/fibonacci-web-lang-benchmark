# Fibonacci Web Framework/Platform Shootout

This repository contains web-apps written in different languages and frameworks that get a number as url parameter and calculate the corresponding fibonacci number.

These implementations are used for benchmarking. The original story can be found here:

https://medium.com/@tschundeee/express-vs-flask-vs-go-acc0879c2122

## The Task
I decided to do a little benchmark of some of my favorite web stacks. So I created a web service that calculates the corresponding fibonacci number to a request parameter.

## The Code
A request to for example http://localhost:3000/10 calculates fib(10) and returns the result (55 in this case) as plain text. I think it’s better to have a little computation per each request because it is more realistic than just serving “Hello World”.

