package main

import (
	"fmt"
	"io/ioutil"
)

func load() []byte {
	dat, err := ioutil.ReadFile("./example.txt")
	fmt.Printf("Error: %s\n", err)
	//fmt.Println(dat)
	return dat
}

func main() {
	fmt.Println(string(load()))
	//fmt.Println(load())
}
