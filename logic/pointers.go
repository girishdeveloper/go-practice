package main

import (
	"fmt"
	"log"
	"strings"
)

type User struct {
	ID   uint64
	Name string
	Paid bool
}

func addElement(cities map[string]string) {
	cities["France"] = "Paris"
}

func addCountries(countries []string) []string {
	for _, s := range []string{"India", "Srilanka"} {
		//log.Println(v, s)
		countries = append(countries, s)
	}
	return countries
}

func addCountries2(countriesPtr *[]string) {
	for _, s := range []string{"India", "Srilanka", "Bhutan"} {
		*countriesPtr = append(*countriesPtr, s)
	}
}

func upper(countries []string) []string {
	for i, v := range countries {
		countries[i] = strings.ToUpper(v)
	}
	return countries
}

func main() {
	var ptr *int
	var answer int = 42
	ptr = &answer
	fmt.Printf("answer is %d and ptr is %d at address %d\n", answer, *ptr, ptr)
	user := User{
		ID:   4,
		Name: "Girish",
		Paid: false,
	}
	userPtr := &user
	userDeref := *userPtr
	user.Name = "Girish Madhavan"
	userDeref.Name = "Axilous"
	fmt.Println(user)
	fmt.Println(userDeref)
	fmt.Println("Maps and Channels are already references to internal structure")
	cities := make(map[string]string)
	addElement(cities)
	log.Println(cities)
	fmt.Println("Slice is a collection of elements of same types.")
	fmt.Println("A slice is internally a structure with 3 fields. Namely,")
	fmt.Println("length, capacity and a pointer to an internal array")
	EUcountries := []string{"Bulgaria", "Austria", "Belgium"}
	log.Println(EUcountries)
	log.Println(addCountries(EUcountries))
	log.Println(len(EUcountries), cap(EUcountries), EUcountries)
	log.Println(upper(EUcountries))
	addCountries2(&EUcountries)
	log.Println(EUcountries)
}
