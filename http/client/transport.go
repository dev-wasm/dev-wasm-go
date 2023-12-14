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

func schemeFromString(s string) proxy.WasiHttp0_2_0_rc_2023_11_10_TypesScheme {
	switch s {
	case "http":
		return proxy.WasiHttp0_2_0_rc_2023_11_10_TypesSchemeHttps()
	case "https":
		return proxy.WasiHttp0_2_0_rc_2023_11_10_TypesSchemeHttps()
	default:
		panic(fmt.Sprintf("Unknown scheme: %s", s))
	}
}

func methodFromString(m string) proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethod {
	switch m {
	case "GET":
		return proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodGet()
	case "PUT":
		return proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodPut()
	case "POST":
		return proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodPost()
	case "DELETE":
		return proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodDelete()
	case "OPTIONS":
		return proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodOptions()
	case "PATCH":
		return proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodPatch()
	case "CONNECT":
		return proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodConnect()
	case "TRACE":
		return proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodTrace()
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
	strstr := []proxy.WasiHttp0_2_0_rc_2023_11_10_TypesTuple2FieldKeyFieldValueT{}
	for k, v := range r.Header {
		// TODO: handle multi-headers here.
		strstr = append(strstr, proxy.WasiHttp0_2_0_rc_2023_11_10_TypesTuple2FieldKeyFieldValueT{k, []uint8(v[0])})
	}
	headers := proxy.StaticFieldsFromList(strstr).Unwrap()

	method := methodFromString(r.Method)
	scheme := proxy.Some(schemeFromString(r.URL.Scheme))

	path_with_query := proxy.Some(r.URL.RequestURI())
	authority := proxy.Some(r.URL.Host)

	req := proxy.NewOutgoingRequest(headers)
	req.SetMethod(method)
	req.SetPathWithQuery(path_with_query)
	req.SetScheme(scheme)
	req.SetAuthority(authority)

	if r.Body != nil {
		s := req.Body().Unwrap().Write().Unwrap()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		s.BlockingWriteAndFlush(b).Unwrap()
	}

	var opts proxy.Option[proxy.WasiHttp0_2_0_rc_2023_11_10_TypesRequestOptions]
	opts = proxy.None[proxy.WasiHttp0_2_0_rc_2023_11_10_TypesRequestOptions]()
	res := proxy.WasiHttp0_2_0_rc_2023_11_10_OutgoingHandlerHandle(req, opts).Unwrap()

	resultOption := res.Get()
	if !resultOption.IsSome() {
		log.Fatalf("No result!")
	}
	result := resultOption.Unwrap().Unwrap().Unwrap()

	response := http.Response{
		StatusCode: int(result.Status()),
		Header:     http.Header{},
	}

	responseHeaders := result.Headers()
	entries := responseHeaders.Entries()

	for _, entry := range entries {
		// TODO: handle multiple headers here.
		response.Header[entry.F0] = []string{string(entry.F1)}
	}

	body := result.Consume().Unwrap()
	stream := body.Stream().Unwrap()

	data := stream.Read(64*1024).Unwrap()

	response.Body = bytesReaderCloser{bytes.NewReader(data)}

	proxy.StaticOutgoingRequestDrop(req)
	proxy.StaticIncomingStreamDrop(stream)
	proxy.StaticIncomingResponseDrop(result)

	return &response, nil
}
