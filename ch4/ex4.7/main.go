/*
Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that
represents a UTF-8-encoded string, in place. Can you do it without allocating
new memory?
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界, привіт"
	b := []byte(s)
	fmt.Printf("%q\n", string(b))
	reverseInPlace(b)
	fmt.Printf("%q\n", string(b))
}

func reverseInPlace(b []byte) {
	l := len(b)
	for i, last := 0, l; i < l; {
		r, s := utf8.DecodeRune(b)
		copy(b, b[s:l-i])
		last -= s
		utf8.EncodeRune(b[last:], r)
		i += s
	}
}
