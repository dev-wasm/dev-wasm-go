package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	wasiclient "github.com/dev-wasm/dev-wasm-go/lib/http/client"
	"github.com/dev-wasm/dev-wasm-go/lib/wasi"
)

func printResponse(r *http.Response) {
	fmt.Printf("Status: %d\n", r.StatusCode)
	for k, v := range r.Header {
		fmt.Printf("%s: %s\n", k, v[0])
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Body: \n%s\n", body)
}

type runner struct{}

func(r runner) Run() wasi.Result[struct{}, struct{}] {
	main()
	return wasi.Ok[struct{}, struct{}](r)
}

func init() {
	wasi.SetExportsWasiCli0_2_0_Run(runner{})
}

func main() {
	client := &http.Client{
		Transport: wasiclient.WasiRoundTripper{},
	}
	req, err := http.NewRequest("GET", "https://postman-echo.com/get", nil)
	if err != nil {
		panic(err.Error())
	}
	if req == nil {
		panic("Nil request!")
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	printResponse(res)

	res, err = client.Post("https://postman-echo.com/post", "application/json", wasiclient.BodyReaderCloser([]byte("{\"foo\": \"bar\"}")))
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	printResponse(res)

	res, err = wasiclient.Put(client, "http://postman-echo.com/put", "application/json", wasiclient.BodyReaderCloser([]byte("{\"baz\": \"blah\"}")))
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	printResponse(res)
}
