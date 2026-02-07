package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	incominghandler "github.com/dev-wasm/dev-wasm-go/lib/wasi/http/incoming-handler"
	"github.com/dev-wasm/dev-wasm-go/lib/wasi/http/types"
	"go.bytecodealliance.org/cm"
)

var h = &handler{
	handler: http.DefaultServeMux,
}

func OK[Shape, T, Err any](val cm.Result[Shape, T, Err]) *T {
	return (&val).OK()
}

func OKOrPanic[Shape, T, Err any](val cm.Result[Shape, T, Err]) *T {
	if !val.IsOK() {
		panic(fmt.Sprintf("a value is not OK as expected: %v", val))
	}
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
	if w.code == 0 {
		w.code = http.StatusOK
	}
	return w.body.Write(data)
}

func schemeToString(scheme *types.Scheme) string {
	if scheme.HTTP() {
		return "http"
	}
	if scheme.HTTPS() {
		return "https"
	}
	return *scheme.Other()
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
	if method.Head() {
		return "HEAD"
	}
	if method.Options() {
		return "OPTIONS"
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
	OKOrPanic(out.BlockingWriteAndFlush(cm.ToList([]uint8(msg))))

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

	scheme := Some(req.Scheme())
	authority := Some(req.Authority())
	path := Some(req.PathWithQuery())
	urlString := fmt.Sprintf("%s://%s%s", schemeToString(scheme), *authority, *path)
	method := req.Method()

	requestBody := OK(req.Consume())
	bodyStream := OK(requestBody.Stream())

	var buff []byte
	for {
		readResult := bodyStream.BlockingRead(1024 * 1024)
		if readResult.IsErr() {
			if readResult.Err().Closed() {
				break
			} else {
				h.HandleError("Reading body failed!", req, responseOut)
				return
			}
		} else {
			bytes := readResult.OK().Slice()
			buff = append(buff, bytes...)
		}
	}
	bodyStream.ResourceDrop()
	requestBody.ResourceDrop()

	fields := req.Headers()
	header := http.Header{}
	for _, tuple := range fields.Entries().Slice() {
		header[string(tuple.F0)] = append(header[string(tuple.F0)], string(tuple.F1.Slice()))
	}

	fields.ResourceDrop()
	req.ResourceDrop()

	goReq, err := http.NewRequest(methodToString(method), urlString, bytes.NewBuffer(buff))
	if err != nil {
		h.HandleError(err.Error(), req, responseOut)
		return
	}
	goReq.Header = header
	if length, ok := header["Content-Length"]; ok {
		if contentLength, err := strconv.Atoi(length[0]); err == nil {
			goReq.ContentLength = int64(contentLength)
		}
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
