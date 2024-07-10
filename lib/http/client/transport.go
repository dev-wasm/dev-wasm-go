package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/dev-wasm/dev-wasm-go/lib/wasi"
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

func schemeFromString(s string) wasi.WasiHttp0_2_0_TypesScheme {
	switch s {
	case "http":
		return wasi.WasiHttp0_2_0_TypesSchemeHttps()
	case "https":
		return wasi.WasiHttp0_2_0_TypesSchemeHttps()
	default:
		panic(fmt.Sprintf("Unknown scheme: %s", s))
	}
}

func methodFromString(m string) wasi.WasiHttp0_2_0_TypesMethod {
	switch m {
	case "GET":
		return wasi.WasiHttp0_2_0_TypesMethodGet()
	case "PUT":
		return wasi.WasiHttp0_2_0_TypesMethodPut()
	case "POST":
		return wasi.WasiHttp0_2_0_TypesMethodPost()
	case "DELETE":
		return wasi.WasiHttp0_2_0_TypesMethodDelete()
	case "OPTIONS":
		return wasi.WasiHttp0_2_0_TypesMethodOptions()
	case "PATCH":
		return wasi.WasiHttp0_2_0_TypesMethodPatch()
	case "CONNECT":
		return wasi.WasiHttp0_2_0_TypesMethodConnect()
	case "TRACE":
		return wasi.WasiHttp0_2_0_TypesMethodTrace()
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
	strstr := []wasi.WasiHttp0_2_0_TypesTuple2FieldKeyFieldValueT{}
	for k, v := range r.Header {
		// TODO: handle multi-headers here.
		strstr = append(strstr, wasi.WasiHttp0_2_0_TypesTuple2FieldKeyFieldValueT{k, []uint8(v[0])})
	}
	headers := wasi.StaticFieldsFromList(strstr).Unwrap()

	method := methodFromString(r.Method)
	scheme := wasi.Some(schemeFromString(r.URL.Scheme))

	path_with_query := wasi.Some(r.URL.RequestURI())
	authority := wasi.Some(r.URL.Host)

	req := wasi.NewOutgoingRequest(headers)
	req.SetMethod(method)
	req.SetPathWithQuery(path_with_query)
	req.SetScheme(scheme)
	req.SetAuthority(authority)

	body := req.Body().Unwrap()
	if r.Body != nil {
		s := body.Write().Unwrap()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		s.BlockingWriteAndFlush(b).Unwrap()
		s.Drop()
	}

	var opts wasi.Option[wasi.WasiHttp0_2_0_TypesRequestOptions]
	opts = wasi.None[wasi.WasiHttp0_2_0_TypesRequestOptions]()
	res := wasi.WasiHttp0_2_0_OutgoingHandlerHandle(req, opts).Unwrap()

	wasi.StaticOutgoingBodyFinish(body, wasi.None[wasi.WasiHttp0_2_0_TypesFields]())

	resultOption := res.Get()
	if resultOption.IsSome() {
		return nil, fmt.Errorf("result already taken!")
	}
	poll := res.Subscribe()
	poll.Block()
	resultOption = res.Get()
	result := resultOption.Unwrap().Unwrap().Unwrap()

	poll.Drop()
	res.Drop()

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
	responseHeaders.Drop()
//	wasi.StaticFieldsDrop(responseHeaders)

	responseBody := result.Consume().Unwrap()
	stream := responseBody.Stream().Unwrap()
	inputPoll := stream.Subscribe()

	data := []uint8{}
	for {
		inputPoll.Block()
		dataResult := stream.Read(64 * 1024)
		if dataResult.IsOk() {
			data = append(data, dataResult.Unwrap()...)
		} else if dataResult.UnwrapErr().Kind() == wasi.WasiIo0_2_0_StreamsStreamErrorKindClosed {
			break
		} else {
			return nil, fmt.Errorf("Error reading response stream")
		}
	}

	response.Body = bytesReaderCloser{bytes.NewReader(data)}

	result.Drop()
//	wasi.StaticIncomingResponseDrop(result)
	//wasi.StaticOutgoingRequestDrop(req)
	//
	//wasi.StaticIncomingStreamDrop(stream)
	//

	return &response, nil
}
