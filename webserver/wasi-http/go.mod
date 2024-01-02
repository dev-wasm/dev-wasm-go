module github.com/dev-wasm/dev-wasm-go/webserver/wasi-http

go 1.21.1

require github.com/dev-wasm/dev-wasm-go/wasi v0.0.0 // indirect

replace github.com/dev-wasm/dev-wasm-go/wasi v0.0.0 => ../../wasi

require github.com/dev-wasm/dev-wasm-go/http v0.0.0-20230731220621-d65072e621f8

replace github.com/dev-wasm/dev-wasm-go/http v0.0.0-20230731220621-d65072e621f8 => ../../http
