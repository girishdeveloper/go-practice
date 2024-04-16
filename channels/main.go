package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(`Goroutines can communicate with each others through channels. 
	A channel can be seen as a pipeline of data between two goroutines. 
	This pipeline can only support a specific type.`)
	var received int
	ch1 := make(chan int, 2)
	ch1 <- 42
	ch1 <- 41
	received = <-ch1
	log.Println(received)
	x, ok := <-ch1
	if !ok {
		log.Println("channel is empty or closed")
	}
	log.Println(x)
	close(ch1)
}
