/*
Exercise 4.5: Write an in-place function to eliminate adjacent
duplicates in a []string slice.
*/

package main

import (
	"fmt"
)

func main() {
	a := []string{"aaa", "aaa", "aaa", "bbb", "ccc", "ccc", "ddd", "ddd"}
	a = deduplicate_adjacent(a)
	fmt.Println(a)
	a = []string{}
	a = deduplicate_adjacent(a)
	fmt.Println(a)
}

func deduplicate_adjacent(s []string) []string {
	var count int
	for i := 1; i < len(s)-count; i++ {
		if s[i] == s[i-1] {
			copy(s[i:], s[i+1:])
			s[len(s)-1] = "" // that is not necessary but more clean I guess
			count++
			i--
		}
	}
	return s[:len(s)-count]
}
