/*
Exercise 3.12: Write a function that reports whether two strings are anagrams of
each other, that is, they contain the same letters in a different order.
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Printf("Usage of %s:\n\t%s string1 string2", os.Args[0], os.Args[0])
		return
	}

	if isAnagram(os.Args[1], os.Args[2]) {
		fmt.Printf("%s %s is anagrams\n", os.Args[1], os.Args[2])
	} else {
		fmt.Printf("%s %s is not anagrams\n", os.Args[1], os.Args[2])
	}

}

func isAnagram(sl, sr string) bool {
	if len(sl) != len(sr) {
		return false
	}
	for _, c := range sl {
		if strings.Count(sl, string(c)) != strings.Count(sr, string(c)) {
			return false
		}
	}
	return true
}
