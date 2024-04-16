package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 1)
	select {
	case rec, ok := <-ch:
		if ok {
			log.Println("received %d", rec)
		}
	case rec, ok := <-time.After(time.Second * 3):
		if ok {
			log.Println("operation timedout as %s", rec)
		}
	}
}
