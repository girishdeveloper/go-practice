package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/homepage", homePageHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homePageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Work is going on")
	fmt.Fprintln(writer, "Are you working here?")
}
