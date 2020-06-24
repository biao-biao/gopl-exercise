package main

import (
	"fmt"
)

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var res int
	for i := 0; i < 64; i++ {
		if x&(1<<i) != 0 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(PopCount(255))
}
