package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/dev-wasm/dev-wasm-go/proxy"
)

type bytesReaderCloser struct {
	*bytes.Reader
}

// Close implements io.Closer.Close.
func (bytesReaderCloser) Close() error {
	return nil
}

func schemeFromString(s string) proxy.TypesScheme {
	switch s {
	case "http":
		return proxy.TypesSchemeHttp()
	case "https":
		return proxy.TypesSchemeHttps()
	default:
		panic(fmt.Sprintf("Unknown scheme: %s", s))
	}
}

func methodFromString(m string) proxy.TypesMethod {
	switch m {
	case "GET":
		return proxy.TypesMethodGet()
	case "PUT":
		return proxy.TypesMethodPut()
	case "POST":
		return proxy.TypesMethodPost()
	case "DELETE":
		return proxy.TypesMethodDelete()
	case "OPTIONS":
		return proxy.TypesMethodOptions()
	case "PATCH":
		return proxy.TypesMethodPatch()
	case "CONNECT":
		return proxy.TypesMethodConnect()
	case "TRACE":
		return proxy.TypesMethodTrace()
	default:
		panic(fmt.Sprintf("Unsupported method: %s", m))
	}
}

func Post(uri, contentType string, body io.ReadCloser) (*http.Response, error) {
	u, e := url.Parse(uri)
	if e != nil {
		return nil, e
	}
	req := http.Request{
		Method: "POST",
		URL:    u,
		Body:   body,
		Header: make(http.Header),
	}
	req.Header["Content-type"] = []string{contentType}
	return Request(req)
}

func Put(uri, contentType string, body io.ReadCloser) (*http.Response, error) {
	u, e := url.Parse(uri)
	if e != nil {
		return nil, e
	}
	req := http.Request{
		Method: "PUT",
		URL:    u,
		Body:   body,
		Header: make(http.Header),
	}
	req.Header["Content-type"] = []string{contentType}
	return Request(req)
}

func Get(u string) (*http.Response, error) {
	ur, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	req := http.Request{
		URL:    ur,
		Method: "GET",
	}
	return Request(req)
}

func Request(r http.Request) (*http.Response, error) {
	if r.Header == nil {
		r.Header = http.Header{}
	}
	if _, ok := r.Header["User-agent"]; !ok {
		r.Header["User-agent"] = []string{"WASI-HTTP-Go/0.0.1"}
	}
	strstr := []proxy.TypesTuple2StringStringT{}
	for k, v := range r.Header {
		// TODO: handle multi-headers here.
		strstr = append(strstr, proxy.TypesTuple2StringStringT{k, v[0]})
	}
	headers := proxy.TypesNewFields(strstr)

	method := methodFromString(r.Method)
	scheme := proxy.Some(schemeFromString(r.URL.Scheme))

	path := r.URL.Path
	authority := r.URL.Host
	query := r.URL.RawQuery

	req := proxy.TypesNewOutgoingRequest(method, path, query, scheme, authority, headers)

	if r.Body != nil {
		s := proxy.TypesOutgoingRequestWrite(req).Unwrap()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		proxy.StreamsWrite(s, b).Unwrap()
	}

	var opts proxy.Option[proxy.TypesRequestOptions]
	opts = proxy.None[proxy.TypesRequestOptions]()
	res := proxy.DefaultOutgoingHttpHandle(req, opts)

	resultOption := proxy.TypesFutureIncomingResponseGet(res)
	if !resultOption.IsSome() {
		log.Fatalf("No result!")
	}
	result := resultOption.Unwrap().Unwrap()

	response := http.Response{
		StatusCode: int(proxy.TypesIncomingResponseStatus(result)),
		Header:     http.Header{},
	}

	responseHeaders := proxy.TypesIncomingResponseHeaders(result)
	entries := proxy.TypesFieldsEntries(responseHeaders)

	for _, entry := range entries {
		// TODO: handle multiple headers here.
		response.Header[entry.F0] = []string{entry.F1}
	}

	stream := proxy.TypesIncomingResponseConsume(result).Unwrap()

	data := proxy.StreamsRead(stream, 64*1024).Unwrap()

	response.Body = bytesReaderCloser{bytes.NewReader(data.F0)}

	proxy.TypesDropOutgoingRequest(req)
	proxy.StreamsDropInputStream(stream)
	proxy.TypesDropIncomingResponse(result)

	return &response, nil
}

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

func main() {
	res, err := Get("https://postman-echo.com/get")
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	printResponse(res)

	res, err = Post("https://postman-echo.com/post", "application/json", bytesReaderCloser{bytes.NewReader([]byte("{\"foo\": \"bar\"}"))})
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	printResponse(res)

	res, err = Put("http://postman-echo.com/put", "application/json", bytesReaderCloser{bytes.NewReader([]byte("{\"baz\": \"blah\"}"))})
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	printResponse(res)
}
