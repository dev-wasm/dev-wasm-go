# Experimental HTTP Client Example
*NB*: this example uses an experimental `wasmtime-http` that incorporates a highly
experimental HTTP client library that is not yet part of the WASI specification.
Use at your own risk, things are likely to change drastically in the future.

## Building
```sh
tinygo build -wasm-abi=generic -target=wasi -o main.wasm
```

## Running without allowed hosts
When you first run this client with `wasmtime-http` no hosts are allowed and so you will see an error.

```sh
$ wasmtime-http --wasi-modules=experimental-wasi-http main.wasm
Request error: (Response error: (7))
```

By default the WASI-http runtime blocks all URLs, to open up the URL, use the `WASI_HTTP_ALLOWED_HOSTS` environment variable:

```sh
$ WASI_HTTP_ALLOWED_HOSTS=https://postman-echo.com wasmtime-http --wasi-modules=experimental-wasi-http main.wasm 
Request status: 200
{"args":{},"headers":{"x-forwarded-proto":"https","x-forwarded-port":"443","host":"postman-echo.com","x-amzn-trace-id":"Root=1-63acc681-2d199daa326ef03d7923f2c4","content-length":"0","content-type":"text/html","user-agent":"wasm32-wasi-http","accept":"*/*"},"url":"https://postman-echo.com/get"}
```