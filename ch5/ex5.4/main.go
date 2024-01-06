/*
Exercise 5.4: Extend the visit function so that it extracts other kinds of
links from the document, such as images, scripts, and style sheets.
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

	links, images, scripts, styles := collect(doc)

	fmt.Println("\nLinks:")
	printSlice(links)

	fmt.Println("\nImages:")
	printSlice(images)

	fmt.Println("\nScripts:")
	printSlice(scripts)

	fmt.Println("\nStyles:")
	printSlice(styles)
}

func printSlice(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func getAttr(n *html.Node, key string) string {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val
		}
	}
	return ""
}

func collect(root *html.Node) (links, images, scripts, styles []string) {
	visitor := func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "a":
				if v := getAttr(n, "href"); len(v) > 0 && v != "/" &&
					!strings.HasPrefix(v, "#") {
					links = append(links, v)
				}
			case "link":
				if v := getAttr(n, "href"); len(v) > 0 &&
					getAttr(n, "rel") == "stylesheet" {
					styles = append(styles, v)
				}
			case "img":
				if v := getAttr(n, "src"); len(v) > 0 {
					images = append(images, v)
				}
			case "script":
				if v := getAttr(n, "src"); len(v) > 0 {
					scripts = append(scripts, v)
				}
			}
		}
	}
	visit(visitor, root)
	return links, images, scripts, styles
}

func visit(visitor func(*html.Node), n *html.Node) {
	if n != nil {
		visitor(n)
		visit(visitor, n.NextSibling)
		visit(visitor, n.FirstChild)
	}
}
