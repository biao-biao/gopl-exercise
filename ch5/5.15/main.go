package main

import "fmt"

func main() {
	fmt.Println(max(1, 2, 3)) // 3
	fmt.Println(min(1, 2, 3)) // 1
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no enough arguments")
	}
	m := vals[0]
	for _, val := range vals {
		if m < val {
			m = val
		}
	}
	return m, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("no enough arguments")
	}
	m := vals[0]
	for _, val := range vals {
		if m > val {
			m = val
		}
	}
	return m, nil
}
