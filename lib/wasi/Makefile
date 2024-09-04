.phony: gen clean run

default: wasi-http

wasi-http: ; git clone https://github.com/WebAssembly/wasi-http; cd wasi-http; git checkout v0.2.0 ; cd ../ ; cp wasi.wit wasi-http/wit

gen: wasi-http ; cd ../ ; wit-bindgen-go generate wasi/wasi-http/wit

clean: ; rm -rf wasi-http