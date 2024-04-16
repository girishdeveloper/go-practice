package main

import (
	"fmt"
)

func main() {
	const version string = "1.3.2"
	fmt.Println("version=", version)
	fmt.Println("Types of constants: type is specified in declaration")
	fmt.Println("Typed constants: not type, has default type and no limits")
	const occupancyLimit = 12
	var occupancyLimit1 uint8
	var occupancyLimit2 int64
	var occupancyLimit3 float32
	occupancyLimit1 = occupancyLimit
	occupancyLimit2 = occupancyLimit
	occupancyLimit3 = occupancyLimit
	fmt.Println(occupancyLimit1, occupancyLimit2, occupancyLimit3)
	fmt.Println("An untyped constant has a default type that is defined by the value assigned to it at compilation.")
	fmt.Println("Untyped constant default types are:")
	fmt.Println("bool, rune, int, float64, complex128, string")
	fmt.Println("Untyped constant has no limits.")
	const profit = 9876653312089144298289
	// illegal: statement below
	var profit2 int64 = profit
	fmt.Println(profit2)
}
