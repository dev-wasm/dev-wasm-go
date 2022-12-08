#!/bin/bash
GOROOT=/usr/local/go tinygo build -wasm-abi=generic -target=wasi -o main.wasm main.go