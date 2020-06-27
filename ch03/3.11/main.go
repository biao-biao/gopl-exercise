package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func comma(s string) string {
	sign, integer, fraction := preprocess(s)
	var buf bytes.Buffer
	n := len(integer)
	a := n % 3

	if a == 0 {
		a += 3
	}

	buf.WriteString(integer[:a])

	for a < n {
		buf.WriteString(",")
		buf.WriteString(integer[a : a+3])
		a += 3
	}

	return sign + buf.String() + "." + fraction
}

func preprocess(s string) (string, string, string) {
	var sign, integer, fraction string
	if s[0] == '-' {
		sign = "-"
		s = s[1:]
	}
	if strings.Contains(s, ".") {
		integer = strings.Split(s, ".")[0]
		fraction = strings.Split(s, ".")[1]
	} else {
		integer = s
	}

	return sign, integer, fraction
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}
