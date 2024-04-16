package relations

import (
	"fmt"
	"log"
)

type Human struct {
	name    string
	age     uint32
	country string
}

type Dog struct {
	name  string
	breed string
	age   uint32
}

type Cat struct {
	name  string
	breed string
	age   uint32
}

func (c *Cat) ReceiveAffection(from Human) {
	log.Printf("cat named %s receives affection from %s\n", c.name, from.name)
}

func (c *Cat) GiveAffection(to Human) {
	log.Printf("cat named %s gives affection to %s\n", c.name, to.name)
}

func (d *Dog) ReceiveAffection(from Human) {
	log.Printf("dog named %s receives affection from %s\n", d.name, from.name)
}

func (d *Dog) GiveAffection(to Human) {
	log.Printf("dog named %s gives affection to %s\n", d.name, to.name)
}

type Stringer interface {
	String() string
}

type DomesticAnimal struct {
	ReceiveAffection (Human)
	GiveAffection    (Human)
	Stringer
}

func (h Human) String() string {
	return fmt.Sprintf("%s is living in %s", h.name, h.country)
}

func SampleHuman() Human {
	return Human{name: "Girish", age: 42, country: "India"}
}

func SampleDog() *Dog {
	return &Dog{name: "Sheru", breed: "Bully Kutta", age: 3}
}

func SampleCat() *Cat {
	return &Cat{name: "Mickey", breed: "Scotish Fold", age: 2}
}
