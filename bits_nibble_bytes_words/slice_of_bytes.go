package main

import (
	"fmt"
)

func main() {
	var slice = make([]byte, 0)
	slice = append(slice, 255)
	slice = append(slice, 10)
	fmt.Println("slice of bytes", slice)
}
