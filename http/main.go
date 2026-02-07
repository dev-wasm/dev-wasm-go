package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"

	wasiclient "github.com/dev-wasm/dev-wasm-go/lib/http/client"
	"github.com/dev-wasm/dev-wasm-go/lib/wasi/cli/run"
	"go.bytecodealliance.org/cm"
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
	// We need at most (align - 1) extra bytes to ensure proper alignment
	newBuf := make([]byte, newSize+align-1)
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

func printResponse(r *http.Response) {
	fmt.Printf("Status: %d\n", r.StatusCode)
	for k, v := range r.Header {
		fmt.Printf("%s: %s\n", k, v[0])
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Body: \n%s\n", body)
}

func Run() cm.BoolResult {
	main()
	return cm.BoolResult(false)
}

func init() {
	run.Exports.Run = Run
}

func main() {
	client := &http.Client{
		Transport: wasiclient.WasiRoundTripper{},
	}
	req, err := http.NewRequest("GET", "https://postman-echo.com/get", nil)
	if err != nil {
		panic(err.Error())
	}
	if req == nil {
		panic("Nil request!")
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	printResponse(res)

	res, err = client.Post("https://postman-echo.com/post", "application/json", wasiclient.BodyReaderCloser([]byte("{\"foo\": \"bar\"}")))
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	printResponse(res)

	res, err = wasiclient.Put(client, "http://postman-echo.com/put", "application/json", wasiclient.BodyReaderCloser([]byte("{\"baz\": \"blah\"}")))
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	printResponse(res)
}
