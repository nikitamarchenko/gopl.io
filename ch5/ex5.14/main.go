/*
Exercise 5.14: Use the breadthFirst function to explore a different structure.
For example, you could use the course dependencies from the topoSort example
(a directed graph), the file system hierarchy on your computer (a tree), or a
list of bus or subway routes downloaded from your city government’s web site
(an undirected graph).
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// !+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(item string) []string {

	var result []string

	info, err := os.Stat(item)

	if err != nil {
		return result
	}

	if info.IsDir() {
		content, err := os.ReadDir(item)
		if err != err {
			return result
		}
		for _, f := range content {
			if f.IsDir() {
				result = append(result, filepath.Join(item, f.Name()))
			} else {
				fmt.Println(filepath.Join(item, f.Name()))
			}
		}
	}

	return result
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
