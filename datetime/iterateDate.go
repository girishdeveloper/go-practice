package main

import (
	"fmt"
	"time"
)

func main() {
	start, err := time.Parse("2006-01-02", "2020-04-13")
	if err != nil {
		panic(err)
	}
	end, err1 := time.Parse("2006-01-02", (time.Now().Format(time.RFC3339))[0:10])
	if err1 != nil {
		panic(err1)
	}
	for i := start; i.Unix() < end.Unix(); i = i.AddDate(1, 0, 0) {
		fmt.Printf("iterate %s\n", i.Format(time.RFC3339))
	}
}
