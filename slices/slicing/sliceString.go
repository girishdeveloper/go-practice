package main

import (
	"fmt"
)

func main() {
	hotelName := "Girish Motel"
	s := hotelName[0:5]
	fmt.Println(s)
	hotelName = "Koresh Motel"
	fmt.Println(s)
}
