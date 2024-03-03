.phony: clean run

default: main_0_2_0.component.wasm

main.wasm: ; tinygo build -o main.wasm -target=wasi main.go

main_0_2_0.embed.wasm: main.wasm ; wasm-tools component embed ../../lib/wasi/wasi-http/wit -w wasi main.wasm -o main_0_2_0.embed.wasm

main_0_2_0.component.wasm: main_0_2_0.embed.wasm ; wasm-tools component new main_0_2_0.embed.wasm -o main_0_2_0.component.wasm --adapt ../../lib/wasi_snapshot_preview1.reactor.0_2_0.wasm

clean: ; rm -rf main*.wasm

run: main_0_2_0.component.wasm ; WASMTIME_BACKTRACE_DETAILS=1 wasmtime serve -Scommon -Dlogging=y main_0_2_0.component.wasm