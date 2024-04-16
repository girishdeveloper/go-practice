package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var hotelName string = "The Gopher Hotel"
	fmt.Println("Hotel", hotelName)
	var roomsAvailable int
	rand.Seed(time.Now().UTC().UnixNano())
	var totalRooms, occupiedRooms = 100, rand.Intn(100)
	roomsAvailable = totalRooms - occupiedRooms
	fmt.Println("Total Rooms", totalRooms)
	fmt.Println("Occupied Rooms", occupiedRooms)
	if roomsAvailable == 0 {
		fmt.Println("No rooms")
	} else {
		fmt.Println("Available rooms", roomsAvailable)
	}
}
