package main

import (
	"fmt"
	"net/http"
)

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