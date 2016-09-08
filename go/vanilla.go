package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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
	n := strings.Trim(r.URL.Path, "/")
	number, err := strconv.Atoi(n)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, fib(number))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":4000", nil)
	fmt.Println("Fin Bench running on Port 4000")
}
