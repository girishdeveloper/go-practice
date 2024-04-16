package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(`A select statement is similar to a switch case statement but for communication operations.

	In a select statement you have cases and an optional default case.
	
	The first non-blocking case will be chosen
	
	If 2 or more cases are not blocking a single one is chosen via an “uniform pseudo-random” selection.
	
	If all cases are blocking, then the default case is chosen`)
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "test"
	select {
	case rec, ok := <-ch1:
		if ok {
			log.Printf("ch1 received value %s", rec)
		}
	case rec, ok := <-ch2:
		if ok {
			log.Printf("ch2 received value %s", rec)
		}
	}
	log.Println("End")
}
