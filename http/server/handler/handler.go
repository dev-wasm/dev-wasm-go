package handler

import (
	"bytes"
	"net/http"

	"github.com/dev-wasm/dev-wasm-go/http/proxy"
)

var h = &handler{
	handler : http.DefaultServeMux,
}

func init() {
	proxy.SetExportsWasiHttp0_2_0_rc_2023_11_10_IncomingHandler(h)
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

func methodToString(method proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethod) string {
	switch method.Kind() {
	case proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodKindGet:
		return "GET"
	case proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodKindPut:
		return "PUT"
	case proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodKindPost:
		return "POST"
	case proxy.WasiHttp0_2_0_rc_2023_11_10_TypesMethodKindDelete:
		return "DELETE"
	default:
		panic("unsupported method")
	}
}

func (h* handler) HandleError(msg string, req proxy.WasiHttp0_2_0_rc_2023_11_10_TypesIncomingRequest, responseOut proxy.WasiHttp0_2_0_rc_2023_11_10_TypesResponseOutparam) {
	hdrs := proxy.NewFields()
	response := proxy.NewOutgoingResponse(hdrs)
	response.SetStatusCode(500)
	body := response.Body().Unwrap()
	resResult := proxy.Ok[proxy.WasiHttp0_2_0_rc_2023_11_10_TypesOutgoingResponse, proxy.WasiHttp0_2_0_rc_2023_11_10_TypesErrorCode](response)
	proxy.StaticResponseOutparamSet(responseOut, resResult)

	out := body.Write().Unwrap()
	out.BlockingWriteAndFlush([]uint8(msg)).Unwrap()
	proxy.StaticOutgoingBodyFinish(body, proxy.None[proxy.WasiHttp0_2_0_rc_2023_11_10_TypesTrailers]())
}

func (h *handler) Handle(req proxy.WasiHttp0_2_0_rc_2023_11_10_TypesIncomingRequest, responseOut proxy.WasiHttp0_2_0_rc_2023_11_10_TypesResponseOutparam) {
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

	headers := []proxy.WasiHttp0_2_0_rc_2023_11_10_TypesTuple2FieldKeyFieldValueT{}
	for key, val := range goRes.header {
		for ix := range val {
			headers = append(headers, proxy.WasiHttp0_2_0_rc_2023_11_10_TypesTuple2FieldKeyFieldValueT{
				F0: key,
				F1: []uint8(val[ix]),
			})
		}
	}
	f := proxy.StaticFieldsFromList(headers).Unwrap()

	res := proxy.NewOutgoingResponse(f)
	res.SetStatusCode(uint16(goRes.code))
	body := res.Body().Unwrap()

	result := proxy.Ok[proxy.WasiHttp0_2_0_rc_2023_11_10_TypesOutgoingResponse, proxy.WasiHttp0_2_0_rc_2023_11_10_TypesErrorCode](res)

	proxy.StaticResponseOutparamSet(responseOut, result)

	stream := body.Write().Unwrap()
	stream.BlockingWriteAndFlush([]byte(goRes.body.Bytes()))
	proxy.StaticOutgoingStreamDrop(stream)
	
	proxy.StaticOutgoingBodyFinish(body, proxy.None[proxy.WasiHttp0_2_0_rc_2023_11_10_TypesTrailers]())
}

func ListenAndServe(handler http.Handler) error {
	if handler != nil {
		h.handler = handler
	}
	return nil
}
