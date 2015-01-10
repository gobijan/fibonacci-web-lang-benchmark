package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// fib returns a function that returns
// successive Fibonacci numbers.
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	stringNumber := p.ByName("number")
	number, _ := strconv.Atoi(stringNumber)
	fmt.Fprintln(w, fib(number))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1) // one core for wrk
	r := httprouter.New()
	r.GET("/:number", indexHandler)
	http.ListenAndServe(":4000", r)
}
