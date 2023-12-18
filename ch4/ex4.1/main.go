/*
Write a function that counts the number of bits that are different in two SHA256
hashes. (See PopCount from Section 2.6.2.)
*/

package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Printf("\tusage %s string1 string2\n", os.Args[0])
		return
	}

	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))

	var diff uint
	for i := 0; i < len(c1); i++ {

		for ii := uint(0); ii < 8; ii++ {
			if c1[i]&(1<<ii) != c2[i]&(1<<ii) {
				diff++
			}
		}
	}

	fmt.Printf("s1=%x\ns2=%x\nidentical=%t\ndiff_bits_count=%d\n",
		c1, c2, c1 == c2, diff)
}
