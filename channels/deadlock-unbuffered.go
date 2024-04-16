package main

import (
	"log"
)

func main() {
	ch := make(chan int)
	go dummy(ch)
	ch <- 45
}

func dummy(ch chan int) {
	log.Println("receiving the message from channel", ch)
}
