/*
Exercise 7.10: The sort.Interface type can be adapted to other uses. Write a
function IsPalindrome(s sort.Interface) bool that reports whether the sequence
s is a palindrome, in other words, reversing the sequence would not change it.
Assume that the elements at indices i and j are equal
if !s.Less(i, j) && !s.Less(j, i).
*/

package palindrome

import (
	"sort"
	"unicode/utf8"
)

type PalindromeString string

func (ps PalindromeString) Len() int {
	return utf8.RuneCountInString(string(ps))
}

func (ps PalindromeString) Less(i, j int) bool {
	r := []rune(ps)
	return r[i] < r[j]
}

func (ps PalindromeString) Swap(i, j int) {
	panic("using PalindromeString for Swap is not allowed")
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < s.Len()/2; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		}
		return false
	}
	return true
}
