package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	outgoinghandler "github.com/dev-wasm/dev-wasm-go/lib/wasi/http/outgoing-handler"
	"github.com/dev-wasm/dev-wasm-go/lib/wasi/http/types"
	"github.com/ydnar/wasm-tools-go/cm"
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

func schemeFromString(s string) types.Scheme {
	switch s {
	case "http":
		return types.SchemeHTTP()
	case "https":
		return types.SchemeHTTPS()
	default:
		panic(fmt.Sprintf("Unknown scheme: %s", s))
	}
}

func methodFromString(m string) types.Method {
	switch m {
	case "GET":
		return types.MethodGet()
	case "PUT":
		return types.MethodPut()
	case "POST":
		return types.MethodPost()
	case "DELETE":
		return types.MethodDelete()
	case "OPTIONS":
		return types.MethodOptions()
	case "PATCH":
		return types.MethodPatch()
	case "CONNECT":
		return types.MethodConnect()
	case "TRACE":
		return types.MethodTrace()
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
		r.Header["User-agent"] = []string{"WASI-HTTP-Go/0.0.2"}
	}
	strstr := []cm.Tuple[types.FieldKey, types.FieldValue]{}
	for k, v := range r.Header {
		// TODO: handle multi-headers here.
		strstr = append(strstr, cm.Tuple[types.FieldKey, types.FieldValue]{types.FieldKey(k), types.FieldValue(cm.ToList([]uint8(v[0])))})
	}
	res := types.FieldsFromList(cm.ToList(strstr))
	headers := res.OK()

	method := methodFromString(r.Method)
	scheme := cm.Some(schemeFromString(r.URL.Scheme))

	path_with_query := cm.Some(r.URL.RequestURI())
	authority := cm.Some(r.URL.Host)

	req := types.NewOutgoingRequest(*headers)
	req.SetMethod(method)
	req.SetPathWithQuery(path_with_query)
	req.SetScheme(scheme)
	req.SetAuthority(authority)

	bodyRes := req.Body()
	body := bodyRes.OK()
	if r.Body != nil {
		writeRes := body.Write()
		s := writeRes.OK()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		s.BlockingWriteAndFlush(cm.ToList([]uint8(b)))
		s.ResourceDrop()
	}

	var opts cm.Option[types.RequestOptions]
	opts = cm.None[types.RequestOptions]()
	hRes := outgoinghandler.Handle(req, opts)
	if !hRes.IsOK() {
		panic("Failed to call client.")
	}

	types.OutgoingBodyFinish(*body, cm.None[types.Fields]())

	future := hRes.OK()
	resultOption := future.Get()
	if !resultOption.None() {
		return nil, fmt.Errorf("result already taken!")
	}
	poll := future.Subscribe()
	poll.Block()
	resultOption = future.Get()
	result := resultOption.Some().OK().OK()

	poll.ResourceDrop()
	future.ResourceDrop()

	response := http.Response{
		StatusCode: int(result.Status()),
		Header:     http.Header{},
	}

	responseHeaders := result.Headers()
	entries := responseHeaders.Entries()

	for _, entry := range entries.Slice() {
		// TODO: handle multiple headers here.
		response.Header[string(entry.F0)] = []string{string(entry.F1.Slice())}
	}
	responseHeaders.ResourceDrop()
	//	wasi.StaticFieldsDrop(responseHeaders)

	bRes := result.Consume()
	responseBody := bRes.OK()
	sRes := responseBody.Stream()
	stream := sRes.OK()
	inputPoll := stream.Subscribe()

	data := []uint8{}
	for {
		inputPoll.Block()
		dataResult := stream.Read(64 * 1024)
		if dataResult.IsOK() {
			data = append(data, dataResult.OK().Slice()...)
		} else if dataResult.Err().Closed() {
			break
		} else {
			return nil, fmt.Errorf("Error reading response stream")
		}
	}

	response.Body = bytesReaderCloser{bytes.NewReader(data)}

	result.ResourceDrop()
	//	wasi.StaticIncomingResponseDrop(result)
	//wasi.StaticOutgoingRequestDrop(req)
	//
	//wasi.StaticIncomingStreamDrop(stream)
	//

	return &response, nil
}
