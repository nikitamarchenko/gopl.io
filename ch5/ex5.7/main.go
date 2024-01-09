/*
Exercise 5.7: Develop startElement and endElement into a general HTML
pretty-printer. Print comment nodes, text nodes, and the attributes of each
element (<a href='...'>). Use short forms like <img/> instead of <img></img>
when an element has no children. Write a test to ensure that the output can be
parsed successfully. (See Chapter 11.)
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	result, err := processHtml(resp.Body)
	if err != nil {
		return err
	}

	fmt.Print(result)

	return nil
}

func processHtml(r io.Reader) (string, error) {

	doc, err := html.Parse(r)
	if err != nil {
		return "", err
	}

	b := strings.Builder{}

	forEachNode(doc, &b, nil, startElement, endElement)

	return b.String(), nil
}

type elementHandler func(n *html.Node, b *strings.Builder, depth *int)

func forEachNode(n *html.Node, b *strings.Builder, depth *int,
	pre, post elementHandler) {

	if depth == nil {
		depth = new(int)
	}

	if pre != nil {
		pre(n, b, depth)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, b, depth, pre, post)
	}

	if post != nil {
		post(n, b, depth)
	}
}

func startElement(n *html.Node, b *strings.Builder, depth *int) {
	if n.Type == html.ElementNode {
		var attrs string
		if len(n.Attr) > 0 {
			b := strings.Builder{}
			for _, attr := range n.Attr {
				b.WriteString(fmt.Sprintf(` %s="%s"`, attr.Key, attr.Val))
			}
			attrs = b.String()
		}

		var closed string
		if n.FirstChild == nil {
			closed = `/`
		}

		b.WriteString(
			fmt.Sprintf("%*s<%s%s%s>\n",
				*depth*2, "", n.Data, attrs, closed))
		*depth++
	}
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if len(text) > 0 {
			b.WriteString(
				fmt.Sprintf("%*s%s\n",
					*depth*2, "", strings.TrimSpace(n.Data)))
		}
	}
	if n.Type == html.CommentNode {
		b.WriteString(
			fmt.Sprintf("%*s<!--%s-->\n",
				*depth*2, "", n.Data))
	}
}

func endElement(n *html.Node, b *strings.Builder, depth *int) {
	if n.Type == html.ElementNode {
		*depth--
		if n.FirstChild != nil {
			b.WriteString(fmt.Sprintf("%*s</%s>\n", *depth*2, "", n.Data))
		}
	}
}
