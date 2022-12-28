package main

import (
	"fmt"
)

type HttpError int32
type ResponseHandle int32

const (
	SUCCESS HttpError = 0
	INVALID_HANDLE HttpError = 1
    MEMORY_NOT_FOUND HttpError = 2
    MEMORY_ACCESS_ERROR HttpError = 3
    BUFFER_TOO_SMALL HttpError = 4
    HEADER_NOT_FOUND HttpError = 5
    UTF_8_ERROR HttpError = 6
    DESTINATION_NOT_ALLOWED HttpError = 7
    INVALID_METHOD HttpError = 8
    INVALID_ENCODING HttpError = 9
    INVALID_URL HttpError = 10
    REQUEST_ERROR HttpError = 11
    RUNTIME_ERROR HttpError = 12
    TOO_MANY_SESSIONS HttpError = 13
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
func req(url *byte, url_len int32, method *byte, method_len int32, headers *byte, headers_len int32, body *byte, body_len int32, status_code *uint16, response_handle *ResponseHandle) HttpError

//go:wasm-module wasi_experimental_http
//export close
func close(handle ResponseHandle) HttpError

//go:wasm-module wasi_experimental_http
//export body_read
func readBody(handle ResponseHandle, buffer *byte, buffer_length int32, bytes_written *uint32) HttpError

//go:wasm-module wasi_experimental_http
//export header_get
func getHeader(handle ResponseHandle, header_name *byte, header_name_len int32, value_buffer *byte, value_buffer_len int32, bytes_written *uint32) HttpError

//go:wasm-module wasi_experimental_http
//export headers_get_all
func getAllHeaders(handle ResponseHandle, value_buffer *byte, value_buffer_len int32, bytes_written *uint32) HttpError

func Request(url string, method string, headers string, body string) (Response, error) {
	cUrl := cstring(url)
	cMethod := cstring(method)
	cHeaders := cstring(headers)
	cBody := cstring(body)

	var statusCode uint16
	var handle ResponseHandle
	httpError := req(&cUrl[0], int32(len(url)), &cMethod[0], int32(len(method)), &cHeaders[0], int32(len(headers)), &cBody[0], int32(len(body)), &statusCode, &handle)
	if httpError != SUCCESS {
		return Response{}, fmt.Errorf("Response error: (%d)", httpError)
	}
	return Response{ handle, int(statusCode) }, nil
}

// TODO: make this closer to golang Response
type Response struct {
	handle ResponseHandle
	// TODO cache return values from body and headers?
	StatusCode int
}

func (r *Response) Body() (string, error) {
	body := make([]byte, 1024 * 1024)
	var written uint32

	err := readBody(r.handle, &body[0], int32(len(body)), &written)
	if err != SUCCESS {
		return "", fmt.Errorf("Request failed: (%v)", err)
	}
	bodyBytes := make([]byte, written)
	for i := 0; i < int(written); i++ {
		bodyBytes[i] = body[i]
	}
	return string(bodyBytes), nil
}

func (r *Response) Close() HttpError {
	return close(r.handle)
}