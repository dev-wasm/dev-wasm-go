module github.com/dev-wasm/dev-wasm-go/examples/component_main

require (
	github.com/dev-wasm/dev-wasm-go v0.0.0-20240903232551-3526e9637f2b // indirect
	github.com/dev-wasm/dev-wasm-go/lib v0.0.0-20240903233842-b25e2e499927 // indirect
	github.com/ydnar/wasm-tools-go v0.1.5 // indirect
)

replace github.com/dev-wasm/dev-wasm-go/lib/ v0.0.0 => ../../lib/

go 1.22.0

toolchain go1.23.0
