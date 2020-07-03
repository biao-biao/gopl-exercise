// See page 141.
// and the equivalent sample in python https://www.toptal.com/python/interview-questions

package main

import "fmt"

func main() {
	var arr []int
	for _, f := range m() {
		arr = append(arr, f(2))
	}

	fmt.Println(arr) // [8 8 8 8]
}

func m() (arr []func(int) int) {
	for i := 0; i < 4; i++ {
		// i := i
		f := func(x int) int {
			return i * x
		}
		arr = append(arr, f)
	}

	return arr
}
