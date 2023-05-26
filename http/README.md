# Experimental HTTP Client Example
*NB*: this example uses an experimental `wasmtime-http` that incorporates a highly
experimental HTTP client library that is not yet part of the WASI specification.
Use at your own risk, things are likely to change drastically in the future.

## Building
```sh
tinygo build -wasm-abi=generic -target=wasi -o main.wasm
```

## Running
```sh
wasmtime --wasi-modules=experimental-wasi-http main.wasm
```

## Regenerating the code 
```sh
git clone https://github.com/WebAssembly/wasi-http
cd wasi-http
git checkout 6c6855a
cd ../
wit-bindgen tiny-go wasi-http/wit -w proxy --out-dir proxy
```

Note that there is currently a bug in `wit-bindgen` which prevents this from working
correctly, it needs a small edit to add a `typedef`.

TODO: extract this out into a general purpose library

