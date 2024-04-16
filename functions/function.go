package main

import (
	"fmt"
)

func computePrice(rate float32, nights int) (price float32) {
	price = rate * float32(nights)
	return
}

func main() {
	girishPrice := computePrice(145.90, 3)
	koreshPrice := computePrice(26.32, 10)
	kilminPrice := computePrice(189.21, 2)
	total := girishPrice + koreshPrice + kilminPrice
	fmt.Printf("TOTAL is %f\n", total)
}
