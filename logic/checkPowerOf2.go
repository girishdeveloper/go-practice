package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter a number.")
	fmt.Scanf("%d", &a)
	if (a & (a - 1)) != 0 {
		fmt.Println("number is not a power of 2")
	} else {
		fmt.Println("number is a power of 2")
	}
}
