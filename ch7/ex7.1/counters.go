/*
Exercise 7.1: Using the ideas from ByteCounter, implement counters for words
and for lines. You will find bufio.ScanWords useful.
*/

package counters

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)


type LineCounter int

func (c *LineCounter) Write(p []byte) (r int, err error) {
	
	r = strings.Count(string(p), "\n")
	*c = LineCounter(r)
	return
}

func (c *LineCounter) String() string {
	return fmt.Sprintf("count = %d", *c)
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return int(*c), scanner.Err()
}

func (c *WordCounter) String() string {
	return fmt.Sprintf("count = %d", *c)
}

