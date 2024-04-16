package main

import (
	"fmt"
)

func Printer() func() {
	k := 1
	return func() {
		fmt.Printf("Print n. %d\n", k)
		k++
	}
}

func main() {
	fmt.Println(`An anonymous function defined into another function F 
	can use elements that are not defined in its scope but, in the scope of F.`)
	p := Printer()
	p()
	p()
	p()
	fmt.Println(`Here we :

	define the variable k in the outer function
	
	we return a new anonymous function that makes use of k.
	
	Hence we have created a closure.
	
	Then in the main function, we will assign to p the return value of printer (which is a function). After that, we will use p.
	
	Remember that p is a function, so to execute it, we have just to add parenthesis at the end :
	
	p() // we execute p, the closure
	The type of p is func().`)
}
