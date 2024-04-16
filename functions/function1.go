package main

import (
	"fmt"
	"math/rand"
	"time"
)

func vacantRooms() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(10)
}

func main() {
	fmt.Println("vacant rooms:", vacantRooms())
}
