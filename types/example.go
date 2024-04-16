package main

import "fmt"

func main() {
	var arr [3]uint8
	var myPointer *int8
	var nameDisplayer func(name, firstname string) string
	var roomNumbers []int8
	var score map[string]uint8
	var received chan<- bool
	fmt.Println(arr, myPointer, nameDisplayer, roomNumbers, score, received)
}
