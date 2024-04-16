package main

import (
	"fmt"
)

func main() {
	fill := make([]int, 5, 10)
	take := []int{1, 2, 3, 4, 5}
	take2 := []int{1, 1, 1, 1, 1, 1}
	fmt.Println("fill", fill, "cap of fill", cap(fill))
	fmt.Println("cap of take", cap(take))
	fmt.Println("cap of take2", cap(take2))
	copy(take2, take)
	fmt.Println(take2, "cap of take2", cap(take2))
	fill = append(fill, 43, 52)
	fmt.Println(fill, "cap of fill", cap(fill))
}
