package main

import "fmt"

func dup(s []string) []string {
	if len(s) == 0 {
		return s
	}

	i := 0
	for j := 1; j < len(s); j++ {
		if s[j] != s[i] {
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}

func main() {
	data := []string{"one", "three", "one", "three", "one", "one", "three"}
	fmt.Printf("%q\n", dup(data))
}
