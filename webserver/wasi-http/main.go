package main

import (
	"fmt"
	"io"
	"net/http"
	"unsafe"

	"github.com/dev-wasm/dev-wasm-go/lib/http/server/handler"
)

// This is required for building the module for some reason
// I don't think it should be. I'm probably doing something
// wrong.
//
// fwiw, I think this is likely buggy and either leaks memory
// or has race conditions.
//
//go:wasmexport cabi_realloc
//export cabi_realloc
func wasmexport_cabi_realloc(ptr, oldSize, align, newSize uint32) uint32 {
	if newSize == 0 {
		return align
	}
	arr := make([]uint8, newSize)
	newPtr := unsafe.Pointer(unsafe.SliceData(arr))
	return uint32(uintptr(newPtr))
}

var count = 0

type myHandler struct{}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count++
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf("Hello from WASM! (%d) %s\n", count, r.URL.Path)))
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(body)
	}
	w.Write([]byte("Headers\n"))
	for k, v := range r.Header {
		w.Write([]byte(k + v[0] + "\n"))
	}
	w.Write([]byte("Query\n"))
	for k, v := range r.URL.Query() {
		w.Write([]byte(k + ":" + v[0] + "\n"))
	}
}

func init() {
	// For some reason
	// http.Handle and http.HandleFunc aren't working :(
	handler.ListenAndServe(myHandler{})
}

func main() {}
