# Experimental HTTP Client Example
*NB*: this example uses an experimental `wasmtime-http` that incorporates a highly
experimental HTTP client library that is not yet part of the WASI specification.
Use at your own risk, things are likely to change drastically in the future.

## Building
```sh
tinygo build -target=wasi -o main.wasm -tags purego
```

## Running
```sh
wasmtime --wasi-modules=experimental-wasi-http main.wasm
```

## Regenerating the code 

install `wit-bindgen v0.11.0` from [releases](https://github.com/bytecodealliance/wit-bindgen/releases/tag/wit-bindgen-cli-0.11.0)

```sh
> wit-bindgen --version
wit-bindgen-cli 0.11.0 (2ec8e3e25 2023-08-28)
```

```sh
git clone https://github.com/WebAssembly/wasi-http
cd wasi-http
git checkout cc86a80
cd ../
wit-bindgen tiny-go wasi-http/wit -w proxy --out-dir proxy
```

TODO: extract this out into a general purpose library

