/*
Exercise 7.10: The sort.Interface type can be adapted to other uses. Write a
function IsPalindrome(s sort.Interface) bool that reports whether the sequence
s is a palindrome, in other words, reversing the sequence would not change it.
Assume that the elements at indices i and j are equal
if !s.Less(i, j) && !s.Less(j, i).
*/

package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	tests := []struct {
		name string
		args PalindromeString
		want bool
	}{
		{"civic", "civic", true},
		{"madam", "madam", true},
		{"radar", "radar", true},
		{"deified", "deified", true},
		{"дід", "дід", true},
		{"not a palindrome", "not a palindrome a ton", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindromeSwap(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	p := PalindromeString("some string")
	p.Swap(0, 1)
}
