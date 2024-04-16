package main

import (
	"fmt"
	"log"
	"reflect"
)

type Funky func(string)

func main() {
	myFunc := func() int {
		fmt.Println("I am anonymous function")
		return 42
	}()
	fmt.Println(reflect.TypeOf(myFunc))
	funcValue := func() int {
		fmt.Println("I am a function literal invoked")
		return 42
	}()
	fmt.Println(reflect.TypeOf(funcValue))
	var f Funky
	f = func(s string) {
		log.Printf("Funky %s", s)
	}
	f("Golang")
}
