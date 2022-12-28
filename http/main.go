package main

import (
	"fmt"
)

func main() {
	response, err := Request(
		"https://postman-echo.com/get",
		"GET", 
		"Content-type: text/html\nUser-agent: wasm32-wasi-http", 
		"")
	defer response.Close()

	if err != nil {
		fmt.Printf("Request error: (%v)\n", err)
		return
	}
	fmt.Printf("Request status: %v\n", response.StatusCode)
	body, err := response.Body()
	if err != nil {
		fmt.Printf("Request error: (%v)\n", err)
		return
	}
	fmt.Printf("%v\n", string(body));
}