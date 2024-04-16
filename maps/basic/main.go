package main

import (
	"fmt"
)

type employee struct {
	Name     string
	Age      uint32
	Position string
}

func main() {
	m := make(map[string]int)
	m["Uruguay"] = 2
	m["Spain"] = 1
	m["Brazil"] = 2
	m["Argentina"] = 2
	fmt.Println(m)
	worldCupWinners := map[int]string{
		1930: "Uruguay",
		1934: "Italy",
		1938: "Italy",
		1950: "Uruguay",
	}
	fmt.Println(worldCupWinners)
	employees := make(map[string]employee)
	emp1 := employee{
		Name:     "Girish",
		Age:      32,
		Position: "Director",
	}
	emp2 := employee{
		Name:     "Koresh",
		Age:      42,
		Position: "Managing Director",
	}
	employees["V0001"] = emp1
	employees["V0002"] = emp2
	fmt.Println(employees)
}
