package invoice

import (
	"fmt"

	"maximilien-andile.com/packageInit/rules/currency"
)

var c = func() string {
	fmt.Println("variable c initialized", b)
	return "value of c"
}()

var a = func() string {
	fmt.Println("variable a initialized")
	return "value of a"
}()

var b = func() string {
	fmt.Println("variable b initialized", a)
	return "value of b"
}()

func init() {
	fmt.Println("--inside invoice package--", c)
}

func Print() {
	fmt.Println("INVOICE number x")
	fmt.Println(54, currency.Eurosymbol())
}
