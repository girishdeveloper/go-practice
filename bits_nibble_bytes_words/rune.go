package main

import (
	"fmt"
)

func main() {
	var str string
	str = "എനിക്ക് Golang ഇഷ്ടമാണ്"
	for _, v := range str {
		fmt.Printf("Unicode %U, character %c, binary %b, octal %o, hexadecimal %x, decimal %d\n", v, v, v, v, v, v)
	}
	var aRune rune = 'Z'
	fmt.Printf("Unicode Code point for &#39;%c&#39;: %U\n", aRune, aRune)
	fmt.Println("When you iterate over a string, you will iterate over runes.")
	var b = "hello"
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
}
