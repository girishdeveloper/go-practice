package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Struct type")
	var personalDetail struct {
		Name     string
		Capacity uint8
		Rooms    uint8
		Smoking  bool
	}
	personalDetail.Name = "Girish"
	fmt.Println(personalDetail)
	//embedded datatypes
	type ExchangeRate map[string]float64
	type BirthDate time.Time
	type Country struct {
		Name        string
		CapitalCity string
	}
	type Hotel struct {
		Name     string
		Capacity uint8
		Rooms    uint8
		Smoking  bool
		Country
	}
	france := Country{
		Name:        "France",
		CapitalCity: "Paris",
	}
	belgium := Country{
		"Belgium",
		"Bruxelles",
	}
	fmt.Println(france, belgium)
	//Pointer type
	type Country1 struct {
		Name        string
		CapitalCity string
	}
	type Hotel1 struct {
		Name string
		*Country1
	}
	hotel := Hotel1{
		Name:     "Utthappa",
		Country1: &Country1{Name: "France"},
	}
	fmt.Println("Hotel is in", hotel.Country1.Name)
	hotel2 := Hotel1{
		Name: "Uzhunu",
		Country1: &Country1{
			Name:        "France",
			CapitalCity: "Paris",
		},
	}
	fmt.Println(hotel2, hotel2.Country1.CapitalCity)
}
