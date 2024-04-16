package main

import (
	"fmt"
)

func addGo(s []string) {
	s = append(s, "Greek")
	fmt.Println(s)
	fmt.Println("in function length of s:", len(s))
	fmt.Println("in function capacity of s:", cap(s))
}

func addGoFixed(s *[]string) {
	*s = append(*s, "Dari")
	fmt.Println(*s)
	fmt.Println("in function length of s:", len(*s))
	fmt.Println("in function capacity of s:", cap(*s))
}

func main() {
	languages := []string{"English", "Hindi", "Chinese", "Russian"}
	fmt.Println(languages)
	fmt.Println("length of languages:", len(languages))
	fmt.Println("capacity of languages:", cap(languages))
	addGo(languages)
	fmt.Println(languages)
	fmt.Println("length of languages:", len(languages))
	fmt.Println("capacity of languages:", cap(languages))
	addGo(languages)
	fmt.Println(languages)
	fmt.Println("length of languages:", len(languages))
	fmt.Println("capacity of languages:", cap(languages))
	addGoFixed(&languages)
	fmt.Println(languages)
	fmt.Println("length of languages:", len(languages))
	fmt.Println("capacity of languages:", cap(languages))
}
