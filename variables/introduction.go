package main

import (
	"fmt"
)

func main() {
	fmt.Println(`Reserved words are : 
	break, default, func, interface, select, case, defer, go, 
	map, struct, chan, else, goto, package, switch, const, 
	fallthrough, if, range, type, continue, for, import, return, 
	var`)
	fmt.Println("variables must be a combination of unicode letters, unicode digits and/ or underscore")
	fmt.Println("and no reserved words.")
	fmt.Println("variables must begin with a letter or an underscore")
	fmt.Println("Types in go language.")
	fmt.Println("string: Character strings")
	fmt.Println("uint,uint8,uint16,uint32,uint64: Unsigned integers")
	fmt.Println("int,int8,int16,int32,int64: Signed integers")
	fmt.Println("bool: Booleans true or false")
	fmt.Println("float32, float64: Floating point numbers")
}
