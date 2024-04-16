package main

import (
	"fmt"
)

type HTTPMethod int

const (
	GET     HTTPMethod = 0
	POST    HTTPMethod = 1
	PUT     HTTPMethod = 2
	DELETE  HTTPMethod = 3
	PATCH   HTTPMethod = 4
	HEAD    HTTPMethod = 5
	OPTIONS HTTPMethod = 6
	TRACE   HTTPMethod = 7
	CONNECT HTTPMethod = 8
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
	handle(POST, map[string]string{"ContentType": "application/json"}, "/home-page")
}
