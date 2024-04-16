package main

import (
	"example/iopack"
	"example/relations"
	"fmt"
	"log"
)

func main() {
	fmt.Println(`An interface is a kind of type that defines 
	a set of methods without implementing them.`)
	var r iopack.Reader
	log.Println(r)
	h1 := relations.SampleHuman()
	d1 := relations.SampleDog()
	c1 := relations.SampleCat()
	d1.ReceiveAffection(h1)
	d1.GiveAffection(h1)
	c1.ReceiveAffection(h1)
	c1.GiveAffection(h1)
	log.Println(h1.String())
}
