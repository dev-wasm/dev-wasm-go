package main

import (
	"fmt"
	"io"
	"net/http"
	"unsafe"

	"github.com/dev-wasm/dev-wasm-go/lib/http/server/handler"
)

// cabi_realloc is required by the WASI Component Model for memory management.
// This implementation allocates new memory and copies old data if needed.
// Note: The allocated memory must not be garbage collected, so we keep a
// reference to prevent GC from reclaiming it.
//
//go:wasmexport cabi_realloc
//export cabi_realloc
func wasmexport_cabi_realloc(ptr, oldSize, align, newSize uint32) uint32 {
	// When freeing (newSize == 0), return align as per Canonical ABI spec
	if newSize == 0 {
		return align
	}
	
	// Allocate new memory with extra space for alignment
	// Go's allocator typically returns 8-byte aligned memory
	newBuf := make([]byte, newSize+align)
	rawPtr := uintptr(unsafe.Pointer(&newBuf[0]))
	
	// Align the pointer to the requested alignment
	alignedPtr := (rawPtr + uintptr(align) - 1) &^ (uintptr(align) - 1)
	newPtr := uint32(alignedPtr)
	
	// Copy old data to new location if reallocation
	if ptr != 0 && oldSize > 0 {
		oldBuf := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), oldSize)
		newSlice := unsafe.Slice((*byte)(unsafe.Pointer(alignedPtr)), newSize)
		copySize := oldSize
		if newSize < oldSize {
			copySize = newSize
		}
		copy(newSlice[:copySize], oldBuf[:copySize])
	}
	
	// Keep reference to prevent GC
	allocated = append(allocated, newBuf)
	
	return newPtr
}

// allocated keeps references to allocated buffers to prevent GC
// Note: This will grow unbounded, but WASM components are typically short-lived
// request handlers where this is acceptable. For long-running scenarios, a more
// sophisticated memory management strategy would be needed.
var allocated [][]byte

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
