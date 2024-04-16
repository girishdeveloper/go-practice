package main

import (
	"fmt"
)

func main() {
	fmt.Println("Swap values of 2 variables without using temporary variable.")
	a := 2
	b := 3
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("Swapped values are", a, b)
}
