package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000, 10))
	elapsed := time.Since(start)
	fmt.Println("Time taken for execution is", elapsed)
}
