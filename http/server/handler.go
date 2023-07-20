package server

import (
	"bytes"
	"net/http"

	"github.com/dev-wasm/dev-wasm-go/http/proxy"
)

type handler struct {
	handler http.Handler
}

type wasmResponseWriter struct {
	header http.Header
	code   int
	body   bytes.Buffer
}

func (w *wasmResponseWriter) Header() http.Header {
	return w.header
}

func (w *wasmResponseWriter) WriteHeader(code int) {
	w.code = code
}

func (w *wasmResponseWriter) Write(data []byte) (int, error) {
	return w.body.Write(data)
}

func methodToString(method proxy.TypesMethod) string {
	switch method.Kind() {
	case proxy.TypesMethodKindGet:
		return "GET"
	case proxy.TypesMethodKindPut:
		return "PUT"
	case proxy.TypesMethodKindPost:
		return "POST"
	case proxy.TypesMethodKindDelete:
		return "DELETE"
	default:
		panic("unsupported method")
	}
}

func (h *handler) Handle(req uint32, responseOut uint32) {
	path := proxy.TypesIncomingRequestPath(req)
	method := proxy.TypesIncomingRequestMethod(req)

	goReq, err := http.NewRequest(methodToString(method), path, &bytes.Buffer{})
	if err != nil {
		panic(err.Error())
	}
	goRes := wasmResponseWriter{
		header: http.Header{},
		code:   -1,
		body:   bytes.Buffer{},
	}
	h.handler.ServeHTTP(&goRes, goReq)

	headers := []proxy.TypesTuple2StringStringT{}
	for key, val := range goRes.header {
		for ix := range val {
			headers = append(headers, proxy.TypesTuple2StringStringT{
				F0: key,
				F1: val[ix],
			})
		}
	}

	f := proxy.TypesNewFields(headers)

	res := proxy.TypesNewOutgoingResponse(uint16(goRes.code), f)

	result := proxy.Result[uint32, proxy.TypesError]{
		Kind: proxy.Ok,
		Val:  res,
	}

	proxy.TypesSetResponseOutparam(responseOut, result)

	stream := proxy.TypesOutgoingResponseWrite(res).Unwrap()
	proxy.StreamsWrite(stream, []byte(goRes.body.Bytes()))

	proxy.TypesDropOutgoingResponse(res)
}

func ListenAndServe(h http.Handler) error {
	if h == nil {
		h = http.DefaultServeMux
	}
	proxy.SetHttp(&handler{
		handler: h,
	})
	//	for true {
	//		time.Sleep(time.Second * 60)
	//	}
	return nil
}
