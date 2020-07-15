package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCount int
type LineCount int

func (c *WordCount) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))

	s.Split(bufio.ScanWords)
	for s.Scan() {
		// fmt.Println(s.Text())
		*c++
	}

	return len(p), s.Err()
}

func (l *LineCount) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))

	for s.Scan() {
		// fmt.Println(s.Text())
		*l++
	}

	return len(p), s.Err()
}

func main() {
	var c WordCount
	var l LineCount

	var name = "c"
	fmt.Fprintf(&c, "a b %s", name)
	fmt.Println(c) // 3

	fmt.Fprintf(&l, "a \nb %s", name)
	fmt.Println(l) // 2
}
