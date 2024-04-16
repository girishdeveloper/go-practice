package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	multiply(s, 2)
	fmt.Println(s)
}

func multiply(s []int, fact int) {
	for i := 0; i < len(s); i++ {
		s[i] = s[i] * fact
	}
}
