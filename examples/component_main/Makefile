.phony: clean run

default: main.component.wasm

wasi-cli-0.2.0: ; wget https://github.com/WebAssembly/wasi-cli/archive/refs/tags/v0.2.0.tar.gz ; tar -xzf v0.2.0.tar.gz ; rm v0.2.0.tar.gz

main.wasm: ; tinygo build -target=wasi -o main.wasm -tags purego

main.embed.wasm: main.wasm wasi-cli-0.2.0; wasm-tools component embed wasi-cli-0.2.0/wit/ -w command main.wasm -o main.embed.wasm

main.component.wasm: main.embed.wasm ; wasm-tools component new main.embed.wasm -o main.component.wasm --adapt ../../lib/wasi_snapshot_preview1.reactor.0_2_0.wasm

clean: ; rm -rf main*.wasm wasi-cli-0.2.0 v0.2.0.tar.gz

run: main.component.wasm ; WASMTIME_BACKTRACE_DETAILS=1 wasmtime -S http --wasm component-model main.component.wasm