.phony: clean run

default: main.component.wasm

wasi-http-0.2.0:
	wget https://github.com/WebAssembly/wasi-http/archive/refs/tags/v0.2.0.tar.gz
	tar -xzf v0.2.0.tar.gz ; rm v0.2.0.tar.gz
	cp client.wit wasi-http-0.2.0/wit/client.wit

main.wasm: ; tinygo build -target=wasip1 -o main.wasm -tags purego main.go

main.embed.wasm: main.wasm wasi-http-0.2.0; wasm-tools component embed wasi-http-0.2.0/wit/ -w client main.wasm -o main.embed.wasm

main.component.wasm: main.embed.wasm ; wasm-tools component new main.embed.wasm -o main.component.wasm --adapt ../lib/wasi_snapshot_preview1.reactor.0_2_0.wasm

clean: ; rm -rf main*.wasm wasi-http-0.2.0

run: main.component.wasm ; WASMTIME_BACKTRACE_DETAILS=1 wasmtime -S http --wasm component-model main.component.wasm