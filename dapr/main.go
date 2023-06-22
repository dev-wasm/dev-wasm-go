package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"

	wasiclient "github.com/dev-wasm/dev-wasm-go/dapr/client"
)

func main() {
	handler.HandleRequestFn = handleRequest
}

func printResponse(r *http.Response) {
	fmt.Printf("Status: %d\n", r.StatusCode)
	for k, v := range r.Header {
		fmt.Printf("%s: %s\n", k, v[0])
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Body: \n%s\n", body)
}

func request() {
	client := http.Client{
		Transport: wasiclient.WasiRoundTripper{},
	}
	res, err := client.Get("https://postman-echo.com/get")
	if err != nil {
		fmt.Println(err.Error())
	}
	printResponse(res)
}

// handleRequest serves a static response from the Dapr sidecar.
func handleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	request()

	// Serve a response that shows the invoked request URI.
	resp.Headers().Set("Content-Type", "text/plain")
	resp.Body().WriteString("hello with requests! " + req.GetURI())
	return // skip any downstream middleware, as we wrote a response.
}