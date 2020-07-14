package main

import (
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Len return the number of elements
func (s *IntSet) Len() int {
	var sum int
	for _, sword := range s.words {
		sum += PopCount(sword)
	}
	return sum
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)

	if s.Has(x) {
		s.words[word] -= 1 << bit
	}
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = make([]uint64, len(s.words))
	copy(t.words, s.words)
	return &t
}

func main() {
	var t IntSet
	t.Add(1)
	t.Add(66)
	t.Add(67)
	s := *(t.Copy())
	fmt.Println(s, t) // {[2 12]} {[2 12]}
	s.Add(2)
	fmt.Println(s, t) // {[6 12]} {[2 12]}
}
