package main

import (
	"fmt"
)

func main() {
	fmt.Println("Types of strings literals:")
	fmt.Println("01. Raw string literals: They are defined between backquotes.")
	fmt.Println("forbidden characters are backquotes")
	fmt.Println("discarded characters are Carriage return")
	fmt.Println("02. Interpreted string literals: They are defined between double quotes.")
	fmt.Println("forbidden characters are new lines and unescaped double quotes")
	raw := `spring rain
	working in multiple lines
	watch the output`
	fmt.Println(raw)
	interpreted := "words described in a \"single line\""
	fmt.Println(interpreted)
}
