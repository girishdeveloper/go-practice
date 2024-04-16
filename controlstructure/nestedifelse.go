package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var ageGirish, ageKoresh int = rand.Intn(110), rand.Intn(110)
	fmt.Println("Girish is", ageGirish, "years old")
	fmt.Println("Koresh is", ageKoresh, "years old")
	if ageGirish > ageKoresh {
		fmt.Println("Girish is elder than Koresh.")
	} else {
		if ageGirish < ageKoresh {
			fmt.Println("Koresh is elder than Girish.")
		} else {
			fmt.Println("Both Girish and Koresh have same age.")
		}
	}
}
