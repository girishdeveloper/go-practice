package main

import (
	"fmt"
)

func main() {
	customers := [4]string{"Girish", "Koresh", "Moresh", "Kilshin"}
	customerSlice := customers[0:1]
	fmt.Printf("customer slice: %s\n", customerSlice) // customer including 0 to 1-1=0
	customerSlice2 := customers[2:4]
	fmt.Printf("customer slice: %s\n", customerSlice2) // customer including 2 to 4-1=3
	//Do slicing copy data?
	customers[0] = "Harish"
	fmt.Println(customerSlice) // value from the same memory address
	fmt.Println("Length is number of values in the slice.")
	fmt.Printf("length of customerSlice is %d\n", len(customerSlice))
	fmt.Printf("capacity of customerSlice is %d\n", cap(customerSlice))
	fmt.Printf("length of customerSlice2 is %d\n", len(customerSlice2))
	fmt.Printf("capacity of customerSlice2 is %d\n", cap(customerSlice2))
}
