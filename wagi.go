package main

import (
	"fmt"
	"os"
	"strings"
)

func parseQueryString(query string) map[string]string {
	result := make(map[string]string)
	if len(query) == 0 {
		return result
	}
	parts := strings.Split(query, "&")
	for _, part := range parts {
		pieces := strings.Split(part, "=")
		switch len(pieces) {
		case 0:
			break
		case 1:
			result[pieces[0]] = ""
		default:
			result[pieces[0]] = pieces[1]
		}
	}
	return result
}

func main() {
	fmt.Println("Content-type: text/html")
	fmt.Println()

	query := parseQueryString(os.Getenv("QUERY_STRING"))
	name := "Unknown"
	if val, ok := query["name"]; ok {
		name = val
	}
	fmt.Printf("<html><body><h3>Hello %s!</h3></body></html>\n", name)
}