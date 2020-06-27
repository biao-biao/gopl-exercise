package main

import "fmt"

func areAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aFreq := make(map[rune]int)
	bFreq := make(map[rune]int)

	for _, c := range a {
		aFreq[c]++
	}
	for _, c := range b {
		bFreq[c]++
	}

	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(areAnagram("abc", "cba"))
}
