package main

import (
	"fmt"

	"github.com/dev-wasm/dev-wasm-go/wasi"
)

type runner struct{}

func(r runner) Run() wasi.Result[struct{}, struct{}] {
	main()
	return wasi.Ok[struct{}, struct{}](r)
}

func init() {
	wasi.SetExportsWasiCli0_2_0_rc_2023_11_10_Run(runner{})
}

func main() {
	fmt.Println("Hello world!")
}
