module github.com/dev-wasm/dev-wasm-go/http

go 1.22.0

toolchain go1.22.5

require (
	github.com/dev-wasm/dev-wasm-go/lib v0.0.0
	go.bytecodealliance.org/cm v0.1.1-0.20250218151459-e57ac0139b6f
)

replace github.com/dev-wasm/dev-wasm-go/lib v0.0.0 => ../lib
