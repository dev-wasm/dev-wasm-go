package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"

	wasiclient "github.com/dev-wasm/dev-wasm-go/lib/http/client"
	"github.com/dev-wasm/dev-wasm-go/lib/http/server/handler"
	"github.com/dev-wasm/dev-wasm-go/lib/wasi/cli/run"
	"github.com/ydnar/wasm-tools-go/cm"
)

type myHandler struct{}

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

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
	handler.ListenAndServe(myHandler{})
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
