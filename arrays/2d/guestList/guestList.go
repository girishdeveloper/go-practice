package guestList

import (
	"github.com/Pallinder/go-randomdata"
)

func Generate() [10][2]string {
	guestList := [10][2]string{}
	for i, _ := range guestList {
		guestList[i] = RoomGuest(i)
	}
	return guestList
}
func RoomGuest(roomId int) [2]string {
	guests := [2]string{}
	guests[0] = randomdata.SillyName()
	guests[1] = randomdata.SillyName()
	return guests
}
