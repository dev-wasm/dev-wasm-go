package main

import (
	"fmt"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

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

type RequestCount struct {
	Count int
}

func makeBuffer(count int) io.Reader {
	body := fmt.Sprintf("[{\"key\": \"count\", \"value\": %d}]", count)
	return bytes.NewBuffer([]byte(body))
}

func request() int {
	client := http.Client{
		Transport: wasiclient.WasiRoundTripper{},
	}
	res, err := client.Get("http://127.0.0.1:3500/v1.0/state/inmemory/count")
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	count := 0
	if res.StatusCode == http.StatusOK {
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}
		count, err = strconv.Atoi(string(data))
		if err != nil {
			panic(err.Error())
		}
	}
	count = count + 1
	res, err = client.Post("http://127.0.0.1:3500/v1.0/state/inmemory", "application/json", makeBuffer(count))
	if err != nil {
		fmt.Println(err.Error())
	}
	printResponse(res)
	return count
}

// handleRequest serves a static response from the Dapr sidecar.
func handleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	if uri := req.GetURI(); !strings.HasPrefix(uri, "/wasm") {
		next = true
		return
	}

	count := request()
	// Serve a response that shows the invoked request URI.
	resp.Headers().Set("Content-Type", "text/plain")
	resp.Body().WriteString(fmt.Sprintf("hello with requests! %s %d", req.GetURI(), count))
	return
}