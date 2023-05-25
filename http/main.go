package main

import (
	"fmt"
	"log"

	"github.com/dev-wasm/dev-wasm-go/proxy"
)

func main() {
	strstr := []proxy.TypesTuple2StringStringT{
		{"User-agent", "WASI-HTTP-Go/0.0.1"},
		{"Content-type", "application/json"},
	}
	headers := proxy.TypesNewFields(strstr)

	method := proxy.TypesMethodGet()
	scheme := proxy.Some(proxy.TypesSchemeHttps())

	path := "/get"
	authority := "postman-echo.com"
	query := ""

	req := proxy.TypesNewOutgoingRequest(method, path, query, scheme, authority, headers)
	var opts proxy.Option[proxy.TypesRequestOptions]
	opts = proxy.None[proxy.TypesRequestOptions]()
	res := proxy.DefaultOutgoingHttpHandle(req, opts)

	resultOption := proxy.TypesFutureIncomingResponseGet(res)
	if !resultOption.IsSome() {
		log.Fatalf("No result!")
	}
	result := resultOption.Unwrap().Unwrap()

	code := proxy.TypesIncomingResponseStatus(result)

	fmt.Printf("Status is %v\n", code)

	responseHeaders := proxy.TypesIncomingResponseHeaders(result)
	entries := proxy.TypesFieldsEntries(responseHeaders)

	for _, entry := range entries {
		fmt.Printf("%s: %s\n", entry.F0, entry.F1)
	}

	stream := proxy.TypesIncomingResponseConsume(result).Unwrap()

	data := proxy.StreamsRead(stream, 64*1024).Unwrap()

	fmt.Printf("%s\n", string(data.F0))

	proxy.TypesDropOutgoingRequest(req)
	proxy.StreamsDropInputStream(stream)
	proxy.TypesDropIncomingResponse(result)
}
