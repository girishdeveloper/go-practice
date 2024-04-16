package main

import (
	"fmt"
)

type HTTPMethod int

const (
	GET     HTTPMethod = iota
	POST    HTTPMethod = iota
	PUT     HTTPMethod = iota
	DELETE  HTTPMethod = iota
	PATCH   HTTPMethod = iota
	HEAD    HTTPMethod = iota
	OPTIONS HTTPMethod = iota
	TRACE   HTTPMethod = iota
	CONNECT HTTPMethod = iota
)

func handle(method HTTPMethod, headers map[string]string, uri string) {
	if method == GET {
		fmt.Println("Method is GET")
	} else {
		fmt.Println("Method is", method)
	}
}

func main() {
	handle(GET, map[string]string{"ContentType": "application/json"}, "/home-page")
	handle(CONNECT, map[string]string{"ContentType": "application/json"}, "/home-page")
}
