package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dev-wasm/dev-wasm-go/lib/wasi/cli/run"
	"go.bytecodealliance.org/cm"
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
	fmt.Printf("Time is: %v\n", time.Now())
	fmt.Printf("Random number is %v\n", rand.Int())
}
