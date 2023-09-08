package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/dev-wasm/dev-wasm-go/http/proxy"
)

type bytesReaderCloser struct {
	*bytes.Reader
}

// Close implements io.Closer.Close.
func (bytesReaderCloser) Close() error {
	return nil
}

func BodyReaderCloser(b []byte) io.ReadCloser {
	return bytesReaderCloser{bytes.NewReader(b)}
}

func schemeFromString(s string) proxy.WasiHttpTypesScheme {
	switch s {
	case "http":
		return proxy.WasiHttpTypesSchemeHttp()
	case "https":
		return proxy.WasiHttpTypesSchemeHttps()
	default:
		panic(fmt.Sprintf("Unknown scheme: %s", s))
	}
}

func methodFromString(m string) proxy.WasiHttpTypesMethod {
	switch m {
	case "GET":
		return proxy.WasiHttpTypesMethodGet()
	case "PUT":
		return proxy.WasiHttpTypesMethodPut()
	case "POST":
		return proxy.WasiHttpTypesMethodPost()
	case "DELETE":
		return proxy.WasiHttpTypesMethodDelete()
	case "OPTIONS":
		return proxy.WasiHttpTypesMethodOptions()
	case "PATCH":
		return proxy.WasiHttpTypesMethodPatch()
	case "CONNECT":
		return proxy.WasiHttpTypesMethodConnect()
	case "TRACE":
		return proxy.WasiHttpTypesMethodTrace()
	default:
		panic(fmt.Sprintf("Unsupported method: %s", m))
	}
}

func Put(client *http.Client, uri, contentType string, body io.ReadCloser) (*http.Response, error) {
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
	return client.Do(&req)
}

type WasiRoundTripper struct{}

func (_ WasiRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.Header == nil {
		r.Header = http.Header{}
	}
	if _, ok := r.Header["User-agent"]; !ok {
		r.Header["User-agent"] = []string{"WASI-HTTP-Go/0.0.1"}
	}
	strstr := []proxy.WasiHttpTypesTuple2StringListU8TT{}
	for k, v := range r.Header {
		// TODO: handle multi-headers here.
		strstr = append(strstr, proxy.WasiHttpTypesTuple2StringListU8TT{k, []uint8(v[0])})
	}
	headers := proxy.WasiHttpTypesNewFields(strstr)

	method := methodFromString(r.Method)
	scheme := proxy.Some(schemeFromString(r.URL.Scheme))

	path_with_querty := proxy.Some(r.URL.RequestURI())
	authority := proxy.Some(r.URL.Host)

	req := proxy.WasiHttpTypesNewOutgoingRequest(method, path_with_querty, scheme, authority, headers)

	if r.Body != nil {
		s := proxy.WasiHttpTypesOutgoingRequestWrite(req).Unwrap()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		proxy.WasiIoStreamsWrite(s, b).Unwrap()
	}

	var opts proxy.Option[proxy.WasiHttpTypesRequestOptions]
	opts = proxy.None[proxy.WasiHttpTypesRequestOptions]()
	res := proxy.WasiHttpOutgoingHandlerHandle(req, opts)

	resultOption := proxy.WasiHttpTypesFutureIncomingResponseGet(res)
	if !resultOption.IsSome() {
		log.Fatalf("No result!")
	}
	result := resultOption.Unwrap().Unwrap()

	response := http.Response{
		StatusCode: int(proxy.WasiHttpTypesIncomingResponseStatus(result)),
		Header:     http.Header{},
	}

	responseHeaders := proxy.WasiHttpTypesIncomingResponseHeaders(result)
	entries := proxy.WasiHttpTypesFieldsEntries(responseHeaders)

	for _, entry := range entries {
		// TODO: handle multiple headers here.
		response.Header[entry.F0] = []string{string(entry.F1)}
	}

	stream := proxy.WasiHttpTypesIncomingResponseConsume(result).Unwrap()

	data := proxy.WasiIoStreamsRead(stream, 64*1024).Unwrap()

	response.Body = bytesReaderCloser{bytes.NewReader(data.F0)}

	proxy.WasiHttpTypesDropOutgoingRequest(req)
	proxy.WasiIoStreamsDropInputStream(stream)
	proxy.WasiHttpTypesDropIncomingResponse(result)

	return &response, nil
}
