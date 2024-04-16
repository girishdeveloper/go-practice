package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var bin, oct int
	var deci int
	var hex string
	fmt.Println("Please enter a binary number.")
	fmt.Scanf("%d", &bin)
	deci = getDecimal(bin)
	fmt.Printf("Decimal value is %d\n", deci)
	fmt.Println("Please enter an octal number.")
	fmt.Scanf("%o", &oct)
	fmt.Printf("Decimal value is %d\n", oct)
	fmt.Println("Please enter an hexa-decimal number.")
	fmt.Scanf("%s", &hex)
	resultHD, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decimal value is", resultHD)
	fmt.Println("")
}

func getDecimal(binary int) int {
	var decimal int = 0
	var remainder int
	var index = 0
	for binary != 0 {
		remainder = binary % 10
		binary = binary / 10
		decimal = decimal + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimal
}
