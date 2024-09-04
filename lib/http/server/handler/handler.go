package handler

import (
	"bytes"
	"net/http"

	incominghandler "github.com/dev-wasm/dev-wasm-go/lib/wasi/http/incoming-handler"
	"github.com/dev-wasm/dev-wasm-go/lib/wasi/http/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

var h = &handler{
	handler: http.DefaultServeMux,
}

func OK[Shape, T, Err any](val cm.Result[Shape, T, Err]) *T {
	return (&val).OK()
}

func Some[T any](val cm.Option[T]) *T {
	return (&val).Some()
}

func theHandler(req types.IncomingRequest, res types.ResponseOutparam) {
	h.Handle(req, res)
}

func init() {
	incominghandler.Exports.Handle = theHandler
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

func methodToString(method types.Method) string {
	if method.Get() {
		return "GET"
	}
	if method.Put() {
		return "PUT"
	}
	if method.Post() {
		return "POST"
	}
	if method.Patch() {
		return "PATCH"
	}
	if method.Connect() {
		return "CONNECT"
	}
	if method.Delete() {
		return "DELETE"
	}
	panic("unsupported method")
}

func (h *handler) HandleError(msg string, req types.IncomingRequest, responseOut types.ResponseOutparam) {
	headers := []cm.Tuple[types.FieldKey, types.FieldValue]{}
	hdrs := cm.ToList(headers)
	res := types.FieldsFromList(hdrs)
	response := types.NewOutgoingResponse(*res.OK())
	response.SetStatusCode(500)
	body := OK(response.Body())
	resResult := cm.OK[cm.Result[types.ErrorCodeShape, types.OutgoingResponse, types.ErrorCode]](response)
	types.ResponseOutparamSet(responseOut, resResult)

	out := OK(body.Write())
	// TODO: test response here.
	out.BlockingWriteAndFlush(cm.ToList([]uint8(msg)))

	types.OutgoingBodyFinish(*body, cm.None[types.Fields]())
}

func (h *handler) Handle(req types.IncomingRequest, responseOut types.ResponseOutparam) {
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

	path := Some(req.PathWithQuery())
	method := req.Method()

	goReq, err := http.NewRequest(methodToString(method), *path, &bytes.Buffer{})
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

	headers := []cm.Tuple[types.FieldKey, types.FieldValue]{}
	for key, val := range goRes.header {
		for ix := range val {
			headers = append(headers, cm.Tuple[types.FieldKey, types.FieldValue]{
				F0: types.FieldKey(key),
				F1: types.FieldValue(cm.ToList([]uint8(val[ix]))),
			})
		}
	}
	f := OK(types.FieldsFromList(cm.ToList(headers)))

	res := types.NewOutgoingResponse(*f)
	res.SetStatusCode(types.StatusCode(goRes.code))
	body := OK(res.Body())

	result := cm.OK[cm.Result[types.ErrorCodeShape, types.OutgoingResponse, types.ErrorCode]](res)

	types.ResponseOutparamSet(responseOut, result)

	stream := OK(body.Write())
	stream.BlockingWriteAndFlush(cm.ToList([]byte(goRes.body.Bytes())))
	stream.ResourceDrop()

	types.OutgoingBodyFinish(*body, cm.None[types.Fields]())
}

func ListenAndServe(handler http.Handler) error {
	if handler != nil {
		h.handler = handler
	}
	return nil
}
