package main

import (
	"fmt"
)

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var res int
	for x != 0 {
		// if x == 0 {
		// 	return res
		// }
		x = x & (x - 1)
		res++
	}
	return res
}

func main() {
	fmt.Println(PopCount(255))
}
