/*
Exercise 4.9: Write a program wordfreq to report the frequency of each word in
an input text file. Call input.Split(bufio.ScanWords) before the first call to
Scan to break the input into words instead of lines.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	m := map[string]int{}

	for input.Scan() {
		m[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
		os.Exit(1)
	}

	fmt.Println("word count")
	for k, v := range m {
		fmt.Printf("%s : %d\n", k, v)
	}
}
