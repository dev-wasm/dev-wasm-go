.phony: clean run

default: main.component.wasm

wasi-http-0.2.0:
	wget https://github.com/WebAssembly/wasi-http/archive/refs/tags/v0.2.0.tar.gz
	tar -xzf v0.2.0.tar.gz ; rm v0.2.0.tar.gz

main.wasm: ; tinygo build -o main.wasm -target=wasi main.go

main.embed.wasm: main.wasm wasi-http-0.2.0 ; wasm-tools component embed wasi-http-0.2.0/wit -w proxy main.wasm -o main.embed.wasm

main.component.wasm: main.embed.wasm ; wasm-tools component new main.embed.wasm -o main.component.wasm --adapt ../../lib/wasi_snapshot_preview1.reactor.0_2_0.wasm

clean: ; rm -rf main*.wasm wasi-http-0.2.0

run: main.component.wasm ; WASMTIME_BACKTRACE_DETAILS=1 wasmtime serve -Scommon -Dlogging=y main.component.wasm