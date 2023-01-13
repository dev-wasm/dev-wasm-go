package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func makeRequest() (*http.Request, error) {
	uriString := os.Getenv("REQUEST_SCHEME") + "localhost" + os.Getenv("REQUEST_URI")
	url, err := url.Parse(uriString)
	if err != nil {
		return nil, err
	}
	result := &http.Request{
		Method: os.Getenv("REQUEST_METHOD"),
		URL: url,
	}

	return result, nil
}

type WagiResponseWriter struct {
	headers http.Header
}

func (w *WagiResponseWriter) Header() http.Header {
	return w.headers
}

func (w *WagiResponseWriter) WriteHeader(code int) {
	fmt.Printf("HTTP/1.0 %d status\r\n", code)
	for key, value := range w.headers {
		fmt.Printf("%s: %s\r\n", key, strings.Join(value, ", "))
	}
	fmt.Print("\r\n")
}

func (w *WagiResponseWriter) Write(buffer []byte) (int, error) {
	return fmt.Print(string(buffer))
}

func Serve(fn func(http.ResponseWriter, *http.Request)) {
	req, err := makeRequest()

	if err != nil {
		fmt.Println("Content-type: text/plain")
		fmt.Println("\n")
		fmt.Println(err.Error())
		return
	}

	rw := &WagiResponseWriter{
		headers: make(http.Header),
	}
	fn(rw, req)
}

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-type", "text/html")
	res.WriteHeader(200)

	query := req.URL.Query()
	name := "Unknown"
	if val, ok := query["name"]; ok {
		name = val[0]
	}
	res.Write([]byte(fmt.Sprintf("<html><body><h3>Hello %s!</h3></body></html>\n", name)))
}

func main() {
	Serve(HandleRequest)
}