package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var n = flag.Int("n", 256, "hash method")

	flag.Parse()

	if *n == 512 {
		for _, v := range flag.Args() {
			fmt.Printf("%x\n", sha512.Sum512([]byte(v)))
		}
	} else if *n == 384 {
		for _, v := range flag.Args() {
			fmt.Printf("%x\n", sha512.Sum384([]byte(v)))
		}
	} else {
		for _, v := range flag.Args() {
			fmt.Printf("%x\n", sha256.Sum256([]byte(v)))
		}
	}

}
