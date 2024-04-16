package main

import (
	"fmt"
)

func main() {
	const hotelName string = "Girish Hotel"
	const longitude, latitude = 28.4129209, -78.833238
	var occupancy int = 12
	fmt.Println(hotelName, longitude, latitude, occupancy)
	fmt.Printf("%T, %T, %T, %T\n", hotelName, longitude, latitude, occupancy)
}
