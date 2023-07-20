package main

import (
	"fmt"
	"net/http"

	"github.com/dev-wasm/dev-wasm-go/http/server"
)

var count = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "WASM-Test 0.0.1")
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(200)
		body := fmt.Sprintf("Hello from WASM! (%d)", count)
		count = count + 1
		w.Write([]byte(body))
	})
	server.ListenAndServe(nil)
}
