package main

import (
	"fmt"

	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(isMore, a1)
	result2 := piscine.IsSorted(isMore, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
