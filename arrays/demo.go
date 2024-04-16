package main

import (
	"fmt"
)

const newValue = "working here"

func UpdateArray(newArray *[2]string) {
	newArray[1] = newValue
}

func main() {
	var newArray = [2]string{"Working", "here"}
	UpdateArray(&newArray)
	fmt.Println("new array is ", newArray)
}
