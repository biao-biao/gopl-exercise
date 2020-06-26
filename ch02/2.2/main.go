package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./lenconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		stdin(os.Stdin)
	} else {
		for _, arg := range os.Args[1:] {
			conv(arg)
		}
	}
}

func stdin(f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		conv(input.Text())
	}
}

func conv(arg string) {
	t, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fm: %v\n", err)
		os.Exit(1)
	}

	f := lenconv.Feet(t)
	m := lenconv.Meter(t)

	fmt.Printf("%s = %s, %s = %s\n", f, lenconv.FToM(f), m, lenconv.MToF(m))
}
