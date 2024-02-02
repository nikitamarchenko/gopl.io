/*
Exercise 7.4: The strings.NewReader function returns a value that satisfies
the io.Reader interface (and others) by reading from its argument, a string.
Implement a simple version of NewReader yourself, and use it to make the HTML
parser (§5.2) take input from a string.
*/

package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

type HTMLStringReader string

func (r *HTMLStringReader) Read (p []byte) (n int, err error) {
	if len(p) != 0 {
		n = copy(p, *r)
		*r = HTMLStringReader((*r)[n:])
		if len(*r) == 0 {
			err = io.EOF
		}
	}
	return 
}

func NewReader (s string) io.Reader {
	var r HTMLStringReader = HTMLStringReader(s)
	return &r
}

func main() {
    doc, err := html.Parse(NewReader("<html><body><a href='example.com'>example.com</a></body></html>"))
    
	if err != nil {
		log.Fatal(err)
    }
    for _, link := range visit(nil, doc) {
        fmt.Println(link)
    }
}

func visit(links []string, n *html.Node) []string {
    if n.Type == html.ElementNode && n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                links = append(links, a.Val)
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links, c)
    }
    return links
}