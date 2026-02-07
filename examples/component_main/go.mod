module github.com/dev-wasm/dev-wasm-go/examples/component_main

go 1.23.0

replace github.com/dev-wasm/dev-wasm-go/lib/ v0.0.0 => ../../lib/

require (
	github.com/dev-wasm/dev-wasm-go/lib v0.0.0
	go.bytecodealliance.org/cm v0.1.1-0.20250218151459-e57ac0139b6f
)
