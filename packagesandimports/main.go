package main

import (
	"example/occupancy"
	"example/room"
	"fmt"
)

func init() {
	fmt.Println("launch program initialization")
}

func main() {
	fmt.Println("launch the program!")
	rate := occupancy.Rate(23, 34)
	level := occupancy.Level(rate)
	fmt.Printf("%s %0.2f%%\n", level, rate)
	room.PrintDetails(102, 4, 3)
}
