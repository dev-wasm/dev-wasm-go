module github.com/dev-wasm/dev-wasm-go/http

go 1.22.0

toolchain go1.22.5

require github.com/dev-wasm/dev-wasm-go/lib v0.0.0

require github.com/ydnar/wasm-tools-go v0.1.5 // indirect

replace github.com/dev-wasm/dev-wasm-go/lib v0.0.0 => ../lib
