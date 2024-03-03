package main

import (
	"fmt"
	"net/http"

	"github.com/dev-wasm/dev-wasm-go/http/server/handler"
)

var count = 0

type myHandler struct{}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count++
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Hello from WASM! (%d) %s", count, r.URL.Path)))
}

func init() {
	// For some reason
	// http.Handle and http.HandleFunc aren't working :(
	handler.ListenAndServe(myHandler{})
}

func main() {}
