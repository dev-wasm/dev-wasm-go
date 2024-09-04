package main

import (
	"fmt"
	"net/http"

	"github.com/dev-wasm/dev-wasm-go/lib/http/server/handler"
)

// This is required for building the module for some reason
// I don't think it should be. I'm probably doing something
// wrong.
//
//go:wasmexport cabi_realloc
//export cabi_realloc
func wasmexport_cabi_realloc(a, b, c, d uint32) uint32 {
	return 0
}

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
