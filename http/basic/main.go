package main

import (
	"log"
	"net/http"
	"time"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	toSend := []byte(`<html><head></head><body>This is ok</body></html>`)
	_, err := w.Write(toSend)
	if err != nil {
		log.Printf("err in response: %s", err)
	}
}

func main() {
	const host = "127.0.0.1"
	const port = "8080"
	myServer := &http.Server{
		Addr:         host + ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      &handler{},
	}
	// launch the server
	log.Fatal(myServer.ListenAndServe())
}
