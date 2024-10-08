package main

import (
	"fmt"
	"os"
)

type WasmGo struct {
	Message string
}

func (w *WasmGo) Print() {
	fmt.Println(w.Message)
}

func (w *WasmGo) Write(file string, text string) error {
	return os.WriteFile(file, []byte(text), 0x444)
}

func (w *WasmGo) Copy(from string, to string) error {
	data, err := os.ReadFile(from)
	if err != nil {
		return err
	}
	return os.WriteFile(to, data, 0x444)
}

func main() {
	w := WasmGo{"Hello Go World!"}
	w.Print()
	if err := w.Write("test.txt", "This is a test\n"); err != nil {
		fmt.Printf("Error writing %s\n", err)
		return
	}
	fmt.Println("Wrote file")
	if err := w.Copy("test.txt", "test-2.txt"); err != nil {
		fmt.Printf("Error writing %s\n", err)
		return
	}
	fmt.Println("Copied file")
}
