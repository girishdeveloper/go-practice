package main

import (
	"fmt"
)

func main() {
	var deci, octalValue int
	fmt.Println("Enter a decimal value.")
	fmt.Scanf("%d", &deci)
	fmt.Printf("Converted octal value %O\n", deci)
	octalValue = getOctal(deci)
	fmt.Printf("Evaluated octal value %d", octalValue)
	fmt.Println("")
}

func getOctal(decimal int) int {
	var octal, remainder int = 0, 0
	index := 1
	for decimal != 0 {
		remainder = decimal % 8
		decimal = decimal / 8
		octal = octal + remainder*index
		index = index * 10
	}
	return octal
}
