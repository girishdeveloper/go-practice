package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mux sync.Mutex
var requestCount int

func main() {
	http.HandleFunc("/status", status)
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func status(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	requestCount++
	mux.Unlock()
	fmt.Fprintf(w, "OK - count %d\n", requestCount)
}
