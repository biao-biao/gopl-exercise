package main

import "fmt"

func main() {
	s := [...]int{1, 2, 3, 4, 5}
	rotate(s[:], 7)
	fmt.Println(s) //[3 4 5 1 2]
}

func rotate(s []int, n int) {
	var l []int
	for i := range s {
		l = append(l, s[(n+i)%len(s)])
	}
	copy(s, l)
}
