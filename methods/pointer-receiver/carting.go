package main

import (
	"fmt"
	"log"
)

type Cat struct {
	Name   string
	Age    uint32
	Colour string
}

func (c *Cat) Rename(name string) {
	c.Name = name
}

func (c Cat) RenameV2(name string) {
	c.Name = name
}

func main() {
	cat := &Cat{Name: "Girish", Age: 13, Colour: "White"}
	fmt.Println("Initial Cat")
	log.Println(*cat)
	cat.Rename("Koresh")
	fmt.Println("Updated Cat")
	log.Println(*cat)
	cat.RenameV2("Kilshin")
	fmt.Println("Did you modify my Cat?")
	log.Println(cat)
}
