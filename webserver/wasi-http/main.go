package main

import (
	"fmt"
	"net/http"

	"github.com/dev-wasm/dev-wasm-go/http/server/handler"
)

var count = 0

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Hello from WASM! (%d)", count)))
	})
	handler.ListenAndServe(nil)
}

func main() {}
