package main

import (
	"fmt"
)

func main() {
	var x, y, z, n uint8
	x = 1
	fmt.Printf("%08b\n", x)
	x = 2
	fmt.Printf("%08b\n", x)
	x = 255
	fmt.Printf("%08b\n", x)
	x = 1
	y = 2
	z = x & y
	fmt.Printf("%08b\n", z)
	x = 100
	y = 200
	z = x | y
	fmt.Printf("%08b\n", z)
	x = 200
	y = 50
	z = x ^ y
	fmt.Printf("%08b\n", z)
	x = 100
	y = 50
	z = x &^ y
	fmt.Printf("%08b\n", z)
	x = 100
	n = 3
	z = x << n
	fmt.Printf("%08b\n", z)
	x = 100
	n = 3
	z = x >> n
	fmt.Printf("%08b\n", z)
}
