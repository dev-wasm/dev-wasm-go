package main

import (
	"fmt"

	"github.com/dev-wasm/dev-wasm-go/lib/wasi/cli/run"
	"github.com/ydnar/wasm-tools-go/cm"
)

func Run() cm.BoolResult {
	main()
	return cm.BoolResult(false)
}

func init() {
	run.Exports.Run = Run
}

func main() {
	fmt.Println("Hello world!")
}
