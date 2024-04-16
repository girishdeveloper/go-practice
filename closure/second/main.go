package main

import (
	"fmt"
	"sort"
)

func main() {
	scores := []int{10, 89, 76, 3, 20, 12}
	fmt.Println("initial scores", scores)
	sort.Slice(scores, func(i int, j int) bool { return scores[i] < scores[j] })
	fmt.Println("Ascending scores", scores)
	sort.Slice(scores, func(i int, j int) bool { return scores[i] > scores[j] })
	fmt.Println("Descending scores", scores)
}
