.phony: clean run

default: main_2023_11_10.component.wasm

# main.wasm: ; tinygo build -o main.wasm -target=wasi main_2023_11_10.go
main.wasm: ; tinygo build -o main.wasm -target=wasi main.go

main_2023_11_10.embed.wasm: main.wasm ; wasm-tools component embed ../../http/proxy/wasi-http/wit main.wasm > main_2023_11_10.embed.wasm

main_2023_11_10.component.wasm: main_2023_11_10.embed.wasm ; wasm-tools component new main_2023_11_10.embed.wasm -o main_2023_11_10.component.wasm --adapt ../../lib/wasi_snapshot_preview1.reactor.2023_11_10.wasm

clean: ; rm -rf main*.wasm

run: main_2023_11_10.component.wasm ; WASMTIME_BACKTRACE_DETAILS=1 wasmtime serve main_2023_11_10.component.wasm