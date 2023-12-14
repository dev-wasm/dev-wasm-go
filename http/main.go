package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	wasiclient "github.com/dev-wasm/dev-wasm-go/http/client"
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

//foo export wasi:cli/run@0.2.0-rc-2023-11-10#run

//export exports_wasi_cli_0_2_0_rc_2023_11_10_run_run
func run() {
	main()
}

func main() {
	client := http.Client{
		Transport: wasiclient.WasiRoundTripper{},
	}
	res, err := client.Get("https://postman-echo.com/get")
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

	res, err = wasiclient.Put(&client, "http://postman-echo.com/put", "application/json", wasiclient.BodyReaderCloser([]byte("{\"baz\": \"blah\"}")))
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	printResponse(res)
}
