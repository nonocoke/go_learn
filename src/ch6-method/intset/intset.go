package intset

import (
	"bytes"
	"fmt"
)

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
	return len(s.words)
}

// Remove remove x from the set
func (s *IntSet) Remove(x int) {
	bit := uint(x%64)
	for i, word := range s.words {
		if 1<<bit == word {
			s.words[i] = 0  // Make it zero
		}
	}
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var temp IntSet
	temp.UnionWith(s)
	return &temp
}

func (s*IntSet) AddAll(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				fmt.Println(64*i + j)
				total += 64*i + j
			}
		}
	}
	return total
}

// String return the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}