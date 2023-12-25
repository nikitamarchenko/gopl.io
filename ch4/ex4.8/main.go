/*
Exercise 4.8: Modify charcount to count letters, digits, and so on in their 
Unicode categories, using functions like unicode.IsLetter.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	cat := map[string]func(rune) bool{
		"Control": unicode.IsControl,
		"Digit  ":   unicode.IsDigit,
		"Graphic": unicode.IsGraphic,
		"Letter ":  unicode.IsLetter,
		"Lower  ":   unicode.IsLower,
		"Mark   ":    unicode.IsMark,
		"Number ":  unicode.IsNumber,
		"Print  ":   unicode.IsPrint,
		"Punct  ":   unicode.IsPunct,
		"Space  ":   unicode.IsSpace,
		"Symbol ":  unicode.IsSymbol,
		"Title  ":   unicode.IsTitle,
		"Upper  ":   unicode.IsUpper,
	}

	catCount := map[string]int{}

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		for n, v := range cat {
			if v(r) {
				catCount[n]++
			}
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	fmt.Printf("\ncat\tcount\n")
	for n, v := range catCount {
		if v > 0 {
			fmt.Printf("%s\t%d\n", n, v)
		}
	}
}

//!-
