package main

import "fmt"

// https://www.toptal.com/python/interview-questions
func main() {
	for _, f := range m() {
		fmt.Println(f(2))
	}
	a := []string{"a", "b"}
	fmt.Printf("%T\n", a)
}

func m() (arr []func(int) int) {
	for i := 0; i < 4; i++ {
		f := func(x int) int {
			return i * x
		}
		arr = append(arr, f)
	}

	return arr
}
