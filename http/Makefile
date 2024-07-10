.phony: clean run

default: main_2023_11_10.component.wasm

main.wasm: ; tinygo build -target=wasi -o main.wasm -tags purego main.go

main_2023_11_10.embed.wasm: main.wasm ; wasm-tools component embed ../lib/wasi/wasi-http/wit/ -w wasi main.wasm -o main_2023_11_10.embed.wasm

main_2023_11_10.component.wasm: main_2023_11_10.embed.wasm ; wasm-tools component new main_2023_11_10.embed.wasm -o main_2023_11_10.component.wasm --adapt ../lib/wasi_snapshot_preview1.reactor.0_2_0.wasm

clean: ; rm -rf main*.wasm

run: main_2023_11_10.component.wasm ; WASMTIME_BACKTRACE_DETAILS=1 wasmtime -S http --wasm component-model main_2023_11_10.component.wasm