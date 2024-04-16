package main

import (
	"fmt"
)

func main() {
	var myArray [2]int
	myArray[0] = 503
	myArray[1] = 560
	fmt.Println("myArray values are")
	fmt.Println(myArray)
	mArray := [2]int{143, 154}
	fmt.Println(mArray)
	b := [2]string{"FR", "US"}
	fmt.Println(b)
	c := [2]float32{234.12, 824.23}
	fmt.Println(c)
	d := [2]int{}
	fmt.Println(d)
	fmt.Printf("%v\n", len(c))
	fmt.Printf("%v\n", cap(c))
	fmt.Printf("%T\n", c)
	for index, element := range myArray {
		fmt.Printf("element at %d is %d\n", index, element)
	}
	fmt.Println("ignore the index")
	for _, element := range c {
		fmt.Printf("element is %f\n", element)
	}
	fmt.Println("ignore the value")
	for index, _ := range b {
		fmt.Printf("index is %d\n", index)
	}
	fmt.Println("iterate array and access it using index")
	for i := 0; i < len(b); i++ {
		fmt.Printf("element at %d is %v\n", i, b[i])
	}
	e := [2]int{12, 13}
	f := [2]int{12, 13}
	if e == f {
		fmt.Println("both are equal")
	} else {
		fmt.Println("both are not equal")
	}
}
