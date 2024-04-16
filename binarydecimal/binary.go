package main

import (
	"fmt"
)

func main() {
	var bin []int
	var deci int = 0
	fmt.Println("Please enter a decimal value.")
	fmt.Scanf("%d", &deci)
	fmt.Printf("value %d is of type %T\n", deci, deci)
	fmt.Printf("Converting to binary\n")
	bin = getBinary(deci)
	for i := (len(bin) - 1); i >= 0; i-- {
		fmt.Printf("%d", bin[i])
	}
	fmt.Println()
}
func getBinary(tmp int) []int {
	var binary []int
	for tmp > 0 {
		binary = append(binary, tmp%2)
		tmp = tmp / 2
	}
	return binary
}
