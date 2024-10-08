// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package stderr represents the imported interface "wasi:cli/stderr@0.2.0".
package stderr

import (
	"github.com/dev-wasm/dev-wasm-go/lib/wasi/io/streams"
	"github.com/ydnar/wasm-tools-go/cm"
)

// GetStderr represents the imported function "get-stderr".
//
//	get-stderr: func() -> output-stream
//
//go:nosplit
func GetStderr() (result streams.OutputStream) {
	result0 := wasmimport_GetStderr()
	result = cm.Reinterpret[streams.OutputStream]((uint32)(result0))
	return
}

//go:wasmimport wasi:cli/stderr@0.2.0 get-stderr
//go:noescape
func wasmimport_GetStderr() (result0 uint32)
