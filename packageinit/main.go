package main

import (
	"fmt"

	"maximilien-andile.com/packageInit/rules/invoice"
)

func init() {
	fmt.Println("main executed")
}

func main() {
	fmt.Println("--start program--")
	invoice.Print()
}
