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

func methodToString(method proxy.WasiHttpTypesMethod) string {
	switch method.Kind() {
	case proxy.WasiHttpTypesMethodKindGet:
		return "GET"
	case proxy.WasiHttpTypesMethodKindPut:
		return "PUT"
	case proxy.WasiHttpTypesMethodKindPost:
		return "POST"
	case proxy.WasiHttpTypesMethodKindDelete:
		return "DELETE"
	default:
		panic("unsupported method")
	}
}

func (h *handler) Handle(req uint32, responseOut uint32) {
	pathAndQuery := proxy.WasiHttpTypesIncomingRequestPathWithQuery(req).Unwrap()
	method := proxy.WasiHttpTypesIncomingRequestMethod(req)


	goReq, err := http.NewRequest(methodToString(method), pathAndQuery, &bytes.Buffer{})
	if err != nil {
		panic(err.Error())
	}
	goRes := wasmResponseWriter{
		header: http.Header{},
		code:   -1,
		body:   bytes.Buffer{},
	}
	h.handler.ServeHTTP(&goRes, goReq)

	headers := []proxy.WasiHttpTypesTuple2StringStringT{}
	for key, val := range goRes.header {
		for ix := range val {
			headers = append(headers, proxy.WasiHttpTypesTuple2StringStringT{
				F0: key,
				F1: val[ix],
			})
		}
	}

	f := proxy.WasiHttpTypesNewFields(headers)

	res := proxy.WasiHttpTypesNewOutgoingResponse(uint16(goRes.code), f)

	result := proxy.Result[uint32, proxy.WasiHttpTypesError]{
		Kind: proxy.Ok,
		Val:  res,
	}

	proxy.WasiHttpTypesSetResponseOutparam(responseOut, result)

	stream := proxy.WasiHttpTypesOutgoingResponseWrite(res).Unwrap()
	proxy.WasiIoStreamsWrite(stream, []byte(goRes.body.Bytes()))

	proxy.WasiHttpTypesDropOutgoingResponse(res)
}

func ListenAndServe(h http.Handler) error {
	if h == nil {
		h = http.DefaultServeMux
	}
	proxy.SetExportsWasiHttpIncomingHandler(&handler{
		handler: h,
	})
	//	for true {
	//		time.Sleep(time.Second * 60)
	//	}
	return nil
}
