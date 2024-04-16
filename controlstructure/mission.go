package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const hotelName string = "The Gopher Hotel"
	const totalRooms int = 32
	const firstRoom int = 110
	rand.Seed(time.Now().UTC().UnixNano())
	var roomsOccupied int = rand.Intn(totalRooms)
	var roomsAvailable int
	var occupancyLevel string
	var occupancyRate float32
	fmt.Println("hotel:", hotelName)
	roomsAvailable = totalRooms - roomsOccupied
	occupancyRate = (float32(roomsAvailable) / float32(totalRooms)) * 100
	if occupancyRate > 70.0 {
		occupancyLevel = "High"
	} else if occupancyRate > 50.0 {
		occupancyLevel = "Medium"
	} else {
		occupancyLevel = "Low"
	}
	fmt.Println("			Occupancy level:", occupancyLevel)
	fmt.Printf("			Occupancy rate: %0.2f%%\n", occupancyRate)
	fmt.Println("Rooms available:", roomsAvailable)
	fmt.Println("Total Rooms:", totalRooms)
	if roomsAvailable > 0 {
		for i := 0; i < roomsAvailable; i++ {
			roomNo := firstRoom + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			fmt.Println(roomNo, ":", size, "people/", nights, "nights")
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}
}
