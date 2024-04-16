package main

import (
	"example/cart"
	"log"
)

func main() {
	//How to call receiver functions?
	newCart := cart.Cart{}
	totalPrice, err := newCart.TotalPrice()
	if err != nil {
		log.Printf("impossible to compute the cart: %s", err)
		return
	} else {
		log.Println("Total  Price", totalPrice.Display())
	}
	err = newCart.Lock()
	if err != nil {
		log.Printf("impossible to lock the cart: %s", err)
	}
}
