/*
Exercise 5.3: Write a function to print the contents of all text nodes in an
HTML document tree. Do not descend into <script> or <style> elements, since
their contents are not visible in a web browser.
*/

package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}

	for _, line := range visit(nil, doc) {
		fmt.Println(line)
	}
}

func isVisibleInBrowser(n *html.Node) bool {
	if n.Type != html.TextNode {
		return false
	}

	p := n.Parent

	if p != nil && p.Type == html.ElementNode &&
		(p.Data == "style" || p.Data == "script" || p.Data == "noscript") {
		return false
	}

	return true
}

func visit(lines []string, n *html.Node) []string {
	if n != nil {
		if isVisibleInBrowser(n) {
			text := strings.TrimSpace(n.Data)
			if len(text) > 0 {
				lines = append(lines, text)
			}
		}
		lines = visit(lines, n.FirstChild)
		lines = visit(lines, n.NextSibling)
	}
	return lines
}
