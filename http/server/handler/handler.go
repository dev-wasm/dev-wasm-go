package handler

import (
	"bytes"
	"net/http"

	"github.com/dev-wasm/dev-wasm-go/wasi"
)

var h = &handler{
	handler: http.DefaultServeMux,
}

func init() {
	wasi.SetExportsWasiHttp0_2_0_IncomingHandler(h)
}

func HandleFunc(pattern string, fn http.HandlerFunc) {
	http.DefaultServeMux.HandleFunc(pattern, fn)
}

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

func methodToString(method wasi.WasiHttp0_2_0_TypesMethod) string {
	switch method.Kind() {
	case wasi.WasiHttp0_2_0_TypesMethodKindGet:
		return "GET"
	case wasi.WasiHttp0_2_0_TypesMethodKindPut:
		return "PUT"
	case wasi.WasiHttp0_2_0_TypesMethodKindPost:
		return "POST"
	case wasi.WasiHttp0_2_0_TypesMethodKindDelete:
		return "DELETE"
	default:
		panic("unsupported method")
	}
}

func (h *handler) HandleError(msg string, req wasi.WasiHttp0_2_0_TypesIncomingRequest, responseOut wasi.WasiHttp0_2_0_TypesResponseOutparam) {
	hdrs := wasi.NewFields()
	response := wasi.NewOutgoingResponse(hdrs)
	response.SetStatusCode(500)
	body := response.Body().Unwrap()
	resResult := wasi.Ok[wasi.WasiHttp0_2_0_TypesOutgoingResponse, wasi.WasiHttp0_2_0_TypesErrorCode](response)
	wasi.StaticResponseOutparamSet(responseOut, resResult)

	out := body.Write().Unwrap()
	out.BlockingWriteAndFlush([]uint8(msg)).Unwrap()
	wasi.StaticOutgoingBodyFinish(body, wasi.None[wasi.WasiHttp0_2_0_TypesTrailers]())
}

func (h *handler) Handle(req wasi.WasiHttp0_2_0_TypesIncomingRequest, responseOut wasi.WasiHttp0_2_0_TypesResponseOutparam) {
	defer func() {
		if r := recover(); r != nil {
			msg := "unknown panic"
			switch t := r.(type) {
			case string:
				msg = t
			case error:
				msg = t.Error()
			default:
				// pass
			}
			h.HandleError(msg, req, responseOut)
		}
	}()

	path := req.PathWithQuery().Unwrap()
	method := req.Method()

	goReq, err := http.NewRequest(methodToString(method), path, &bytes.Buffer{})
	if err != nil {
		h.HandleError(err.Error(), req, responseOut)
		return
	}
	goRes := wasmResponseWriter{
		header: http.Header{},
		code:   -1,
		body:   bytes.Buffer{},
	}
	h.handler.ServeHTTP(&goRes, goReq)

	headers := []wasi.WasiHttp0_2_0_TypesTuple2FieldKeyFieldValueT{}
	for key, val := range goRes.header {
		for ix := range val {
			headers = append(headers, wasi.WasiHttp0_2_0_TypesTuple2FieldKeyFieldValueT{
				F0: key,
				F1: []uint8(val[ix]),
			})
		}
	}
	f := wasi.StaticFieldsFromList(headers).Unwrap()

	res := wasi.NewOutgoingResponse(f)
	res.SetStatusCode(uint16(goRes.code))
	body := res.Body().Unwrap()

	result := wasi.Ok[wasi.WasiHttp0_2_0_TypesOutgoingResponse, wasi.WasiHttp0_2_0_TypesErrorCode](res)

	wasi.StaticResponseOutparamSet(responseOut, result)

	stream := body.Write().Unwrap()
	stream.BlockingWriteAndFlush([]byte(goRes.body.Bytes()))
	stream.Drop()

	wasi.StaticOutgoingBodyFinish(body, wasi.None[wasi.WasiHttp0_2_0_TypesTrailers]())
}

func ListenAndServe(handler http.Handler) error {
	if handler != nil {
		h.handler = handler
	}
	return nil
}
