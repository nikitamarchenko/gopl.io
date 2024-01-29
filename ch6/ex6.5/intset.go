/*

Exercise 6.5: The type of each word used by IntSet is uint64, but 64-bit 
arithmetic may be inefficient on a 32-bit platform. Modify the program to use 
the uint type, which is the most efficient unsigned integer type for the 
platform. Instead of dividing by 64, define a constant holding the effective 
size of uint in bits, 32 or 64. You can use the perhaps too-clever 
expression 32 << (^uint(0) >> 63) for this purpose.

*/

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
	"math/bits"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// const uintSize = 32 << (^uint(0) >> 63)
// or
// const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
// or
// bits.UintSize from "math/bits"

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bits.UintSize, uint(x%bits.UintSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bits.UintSize, uint(x%bits.UintSize)
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

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits.UintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bits.UintSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

// pop count from ch2
func PopCountByClearing(x uint) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

// return the number of elements
func (s *IntSet) Len() (r int) {
	for _, w := range s.words {
		r += PopCountByClearing(w)
	}
	return r
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/bits.UintSize, uint(x%bits.UintSize)
	if word < len(s.words) {
		s.words[word] ^= 1 << bit
	}
}

// remove all elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// return a copy of the set
func (s *IntSet) Copy() (r *IntSet) {
	r = &IntSet{}
	r.words = make([]uint, len(s.words))
	if r := copy(r.words, s.words); r != len(s.words) {
		panic("error while copy words slice")
	}
	return r
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) AddAll(values ...int) {
	for _, v := range values {
		s.Add(v)
	}
}

func (s *IntSet) Elems() (r []uint) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bits.UintSize; j++ {
			if word&(1<<uint(j)) != 0 {
				r = append(r, uint(bits.UintSize*i+j))
			}
		}
	}
	return r
}
