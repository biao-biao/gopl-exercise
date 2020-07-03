package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	var diff int
	for i := range c1 {
		diff += PopCount(uint64(c1[i] ^ c2[i]))
	}

	fmt.Println(diff)
}
