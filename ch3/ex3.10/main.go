/*
Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead
of string concatenation.
*/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//

package main

import (
	"bytes"
	"fmt"

	"os"
	"strconv"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {

	for _, c := range s {
		if _, err := strconv.Atoi(string(c)); err != nil {
			fmt.Printf("%s is not a decimal integer\n", s)
			os.Exit(1)
		}
	}

	var buf bytes.Buffer
	l := len(s) - 1

	if len(s) <= 3 {
		return s
	}

	for i := range s {
		buf.WriteByte(s[i])
		if i < l && (l-i)%3 == 0 {
			buf.WriteRune(',')
		}
	}

	return buf.String()
}
