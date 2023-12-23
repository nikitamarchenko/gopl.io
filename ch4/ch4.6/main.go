/*
Exercise 4.6: Write an in-place function that squashes each run of adjacent
Unicode spaces (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a
single ASCII space.
*/

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "  Hello, 世界, \n,  \u0085  ! \u00A0 привіт \f! \v"
	b := []byte(s)
	fmt.Printf("\"%s\"\n", string(b))
	fmt.Printf("%q\n", string(b))
	b = squashAdjacentSpaces(b)
	fmt.Printf("%q\n", string(b))
}

func squashAdjacentSpaces(b []byte) []byte {
	var last rune
	var removed int
	for i := 0; i < len(b)-removed; {
		r, s := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(r) && unicode.IsSpace(last) {
			copy(b[i:], b[i+s:])
			removed += s
			continue
		}
		last = r
		i += s
	}
	return b[:len(b)-removed]
}
