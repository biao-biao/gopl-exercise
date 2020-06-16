package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		counts := make(map[string]int)

		countlines(f, counts)
	}

}

func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	count := 1

	for input.Scan() {
		counts[input.Text()]++
		if len(counts) != count {
			fmt.Println(f.Name())
			break
		}
		count++
	}
}
