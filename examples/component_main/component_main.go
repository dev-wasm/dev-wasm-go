package main

import (
	"fmt"

	"github.com/dev-wasm/dev-wasm-go/lib/wasi"
	"github.com/dev-wasm/dev-wasm-go/lib/wasi/cli/run"
)

type runner struct{}

func (r runner) Run() wasi.Result[struct{}, struct{}] {
	main()
	return wasi.Ok[struct{}, struct{}](r)
}

func init() {
	run.Exports.Run = runner.Run
}

func main() {
	fmt.Println("Hello world!")
}
