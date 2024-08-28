# Devcontainer WASM-Go
Simple devcontainer for Go development

# Usage

## Github Codespaces
Just click the button:

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://github.com/codespaces/new?hide_repo_select=true&ref=main&repo=575629782)



## Visual Studio Code
Note this assumes that you have the VS code support for remote containers and `docker` installed 
on your machine.

```sh
git clone https://github.com/dev-wasm/dev-wasm-go
cd dev-wasm-go
code ./
```

Visual studio should prompt you to see if you want to relaunch the workspace in a container, you do.

# Building and Running

## Simple example
```sh
# Build with TinyGo
tinygo build -target=wasi -o main.wasm main.go
# Or Build with Golang native
GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go

# Run it
wasmtime --dir .::/ main.wasm
```



## WASM Web serving with WASI-HTTP
There is an example of web serving via the [WASI-HTTP](https://github.com/WebAssembly/wasi-http) API
in `webserver/wasi-http`. 

To build:
```sh
# Install the wit interface files, this only needs to be done once.
cd lib/wasi
make wasi-http

# Return to root
cd ../../

# Build the wasm component
cd webserver/wasi-http
make clean
make default
make run
```

Once it is running you can connect to it via http://localhost:8080

## WASM CGI web serving with lighttpd
There is an example of web serving via WebAssembly + CGI (WAGI) in
the `webserver/wagi` directory. It uses the lighttpd web server and `mod_cgi`.
See the `webserver/lighttpd.conf` file for more details.

```sh
tinygo build -target=wasi -o wagi.wasm webserver/wagi.go
lighttpd -D -f webserver/lighttpd.conf
```

Once the server is running, VS Code or Codespaces should prompt you to connect to the open port.

## HTTP Client Example
There is a more complicated example in the [`http` directory](./http/) which shows an example 
of making an HTTP client call using the experimental wasi-http support in `wasmtime`.

