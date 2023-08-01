.phony: gen clean run

default: main.wasm

wasi-http: ; git clone https://github.com/WebAssembly/wasi-http; cd wasi-http; git checkout 1c95bc2

gen: wasi-http ; wit-bindgen tiny-go ${PWD}/wasi-http/wit -w proxy

clean: ; rm -rf wasi-http
