package currency

import (
	"fmt"
)

var f = func() string {
	fmt.Println("--variable f initialized--")
	return "test"
}()

func init() {
	fmt.Println("--currency limit--")
}

func Eurosymbol() string {
	return "EUR"
}
