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
	switch ageSum := ageGirish + ageKoresh; ageSum {
	case 10:
		fmt.Println("ageGirish + ageKoresh = 10")
	case 20, 30, 40:
		fmt.Println("ageGirish + ageKoresh = 20 or 30 or 40")
	case 2 * ageKoresh:
		fmt.Println("ageGirish + ageKoresh = 2 times ageKoresh")
	default:
		fmt.Println("ageGirish + ageKoresh =", ageSum)
	}
}
