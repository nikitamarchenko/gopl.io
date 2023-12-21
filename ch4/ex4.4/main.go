/*
Exercise 4.4: Write a version of rotate that operates in a single pass.
*/

package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	rotate(a)
	fmt.Println(a)
}

func rotate(s []int) {
	if len(s) > 1 {
		first := s[0]
		copy(s, s[1:])
		s[len(s)-1] = first
	}
}
