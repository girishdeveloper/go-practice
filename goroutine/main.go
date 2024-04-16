package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("launch first goroutine")
	go printNumber("first")
	fmt.Println("launch second goroutine")
	go printNumber("second")
	time.Sleep(1 * time.Minute)
}

func printNumber(str string) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		i++
		fmt.Println(i, str)
	}
}
