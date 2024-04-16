package room

import (
	"fmt"
)

func PrintDetails(roomNumber int, size int, nights int) {
	fmt.Printf("%d : %d people/ %d nights\n", roomNumber, size, nights)
}
