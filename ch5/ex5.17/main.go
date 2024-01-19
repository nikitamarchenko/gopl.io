/*
Exercise 5.17: Write a variadic function ElementsByTagName that, given an HTML
node tree and zero or more names, returns all the elements that match one of
those names. Here are two example calls:

Click here to view code image

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node

images := ElementsByTagName(doc, "img")
headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
*/

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}

	for _, link := range ElementsByTagName(doc, "img") {
		print(link)
	}

	for _, link := range ElementsByTagName(doc, "h1", "h2", "h3", "h4") {
		print(link)
	}
}

func print(n *html.Node) {
	
	fmt.Println(n.Data)
	for _, a := range n.Attr {
		fmt.Println("  ", a.Key, a.Val)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	r := []*html.Node{}
	if doc != nil {
		if doc.Type == html.ElementNode {
			for _, n := range name {
				if n == doc.Data {
					r = append(r, doc)
				}
			}
		}
		r = append(r, ElementsByTagName(doc.NextSibling, name...)...)
		r = append(r, ElementsByTagName(doc.FirstChild, name...)...)
	}
	return r
}
