.phony: gen clean run

default: main.wasm

wasi-http: ; git clone https://github.com/WebAssembly/wasi-http; cd wasi-http; git checkout 53e2d25

# Note that this doesn't quite work right there's a bug in wit-bindgen 0.4.0 that missing a C definition
# it is fixed in later versions, but they don't understand this wit syntax *bleeding-edge*
gen: wasi-http ; wit-bindgen tiny-go ${PWD}/wasi-http/wit -w proxy

clean: ; rm -rf wasi-http