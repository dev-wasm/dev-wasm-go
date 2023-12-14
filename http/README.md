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


