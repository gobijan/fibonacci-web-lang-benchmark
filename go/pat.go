package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/bmizerany/pat"
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	stringNumber := r.URL.Query().Get(":number")
	number, _ := strconv.Atoi(stringNumber)
	fmt.Fprintln(w, fib(number))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1) // one core for wrk
	m := pat.New()
	//m.Get("/{number:[0-9]+}", http.HandlerFunc(IndexHandler))
	m.Get("/:number", http.HandlerFunc(indexHandler))
	http.Handle("/", m)
	http.ListenAndServe(":4000", nil)
}
