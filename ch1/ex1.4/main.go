/*
Exercise 1.4: Modify dup2 to print the names of all files in which each
duplicated line occurs.
*/
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]*list.List)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			l := filenames[line]
			for e := l.Front(); e != nil; e = e.Next() {
				fmt.Println("\t", e.Value)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string]*list.List) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		s := input.Text()
		counts[s]++
		l := filenames[s]
		if l == nil {
			l = list.New()
			filenames[s] = l
		}
		l.PushBack(f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
