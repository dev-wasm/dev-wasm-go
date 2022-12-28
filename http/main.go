package main

import (
	"fmt"
)

// Copied from https://github.com/tinygo-org/tinygo/blob/release/src/syscall/syscall_libc.go#L304
// because it's not exported
func cstring(s string) []byte {
	data := make([]byte, len(s)+1)
	copy(data, s)
	// final byte should be zero from the initial allocation
	return data
}

//go:wasm-module wasi_experimental_http
//export req
func req(url *byte, url_len int32, method *byte, method_len int32, headers *byte, headers_len int32, body *byte, body_len int32, status_code *uint16, response_handle *uint32) int32

func request(url string, method string, headers string, body string, handle *uint32) (statusCode uint16, httpError int32) {
	cUrl := cstring(url)
	cMethod := cstring(method)
	cHeaders := cstring(headers)
	cBody := cstring(body)
	httpError = req(&cUrl[0], int32(len(url)), &cMethod[0], int32(len(method)), &cHeaders[0], int32(len(headers)), &cBody[0], int32(len(body)), &statusCode, handle)
	return statusCode, httpError
}

func main() {
	var handle uint32
	_, err := request("https://google.com", "GET", "", "", &handle)
	fmt.Printf("REQUEST: %v\n", result)
}