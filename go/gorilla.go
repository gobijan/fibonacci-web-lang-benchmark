package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	stringNumber := vars["number"]
	number, _ := strconv.Atoi(stringNumber)
	fmt.Fprintln(w, fib(number))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1) // one core for wrk
	r := mux.NewRouter()
	r.HandleFunc("/{number:[0-9]+}", indexHandler)
	http.Handle("/", r)
	http.ListenAndServe(":4000", nil)
}
