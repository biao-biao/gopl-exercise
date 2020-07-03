package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	a := n % 3

	if a == 0 {
		a += 3
	}

	buf.WriteString(s[:a])

	for a < n {
		buf.WriteString(",")
		buf.WriteString(s[a : a+3])
		a += 3
	}

	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}
