package main

import (
	"fmt"
	"time"
)

func main() {
	const ISTTimeZone string = "Asia/Kolkata"
	now := time.Now()
	fmt.Println("Hello World")
	fmt.Println("Time is:", now)
	loc, err := time.LoadLocation(ISTTimeZone)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("location", loc)
	}
}
