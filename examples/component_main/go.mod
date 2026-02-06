module github.com/dev-wasm/dev-wasm-go/examples/component_main

require (
	github.com/dev-wasm/dev-wasm-go/lib v0.0.0-20240903233842-b25e2e499927
	github.com/ydnar/wasm-tools-go v0.1.5
)

replace github.com/dev-wasm/dev-wasm-go/lib/ v0.0.0 => ../../lib/

go 1.22.0

toolchain go1.23.0
