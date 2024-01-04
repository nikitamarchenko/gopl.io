/*
Exercise 5.2: Write a function to populate a mapping from element names—p, div,
span, and so on—to the number of elements with that name in an HTML document
tree.
*/

package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
	"golang.org/x/net/html"
)

func sortKeys(m map[string]int) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}

	result := make(map[string]int, 0)
	visit(result, doc)
	fmt.Println("\n  Name      Count")
	fmt.Println("  ───────────────")
	for _, name := range sortKeys(result) {
		line := fmt.Sprintf("%-10s %4d", name, result[name])
		line = strings.ReplaceAll(line, " ", "·")
		fmt.Printf("  %s\n", line)
	}
	fmt.Println()
}

func visit(count map[string]int, n *html.Node) {
	if n != nil {
		if n.Type == html.ElementNode {
			count[n.Data]++
		}
		visit(count, n.NextSibling)
		visit(count, n.FirstChild)
	}
}
