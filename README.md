# Devcontainer WASM-TS
Simple devcontainer for Go development

# Usage

## Github Codespaces
Just click the button:


## Visual Studio Code
Note this assumes that you have the VS code support for remote containers and `docker` installed 
on your machine.

```sh
git clone https://github.com/brendandburns/dev-wasm-go
cd dev-wasm-go
code ./
```

Visual studio should prompt you to see if you want to relaunch the workspace in a container, you do.

# Building and Running

```sh
./build.sh
wasmtime main.wasm
```
