package main

import (
	"fmt"
	"strconv"
)

func main() {
	var deci int64 = 0
	fmt.Println("Enter a decimal value.")
	fmt.Scanf("%d", &deci)
	fmt.Printf("Converted Hexa-decimal value is %X\n", deci)
	output := strconv.FormatInt(deci, 16)
	fmt.Println("Evaluated Hexa-decimal value is", output)
	fmt.Println("")
}
