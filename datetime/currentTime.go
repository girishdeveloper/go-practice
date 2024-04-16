package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("current time is", time.Now())
	fmt.Println("time Epoch", time.Now().Unix())
	fmt.Println("Formatted time", time.Now().Format(time.RFC3339))
	loc, err := time.LoadLocation("Europe/London")
	if err != nil {
		panic(err)
	}
	fmt.Println("Date in a location", time.Now().In(loc))
	fmt.Println("Parse a date/time contained in a string")
	timeToParse := "2019-02-15T07:33-05:00"
	t, err := time.Parse("2006-01-02T15:04-07:00", timeToParse)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parsed time is", t)
	end := time.Now()
	elapsedTime := end.Sub(start)
	fmt.Printf("process took %s\n", elapsedTime)
}
