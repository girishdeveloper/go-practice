package main

import (
	"fmt"
)

func main() {
	a := make([]int, 3)
	fmt.Println("a is initially", a)
	a[0] = 21
	a[2] = 41
	fmt.Println("a is now", a)
	s2 := []int{10, 12}
	fmt.Println("s2 is now", s2)
}
