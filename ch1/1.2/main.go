package main

import (
	"fmt"
	"os"
)

func main() {
	for i := range os.Args {
		fmt.Println(i, os.Args[i])
	}
}
