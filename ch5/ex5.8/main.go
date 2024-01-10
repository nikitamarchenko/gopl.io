/*
Exercise 5.8: Modify forEachNode so that the pre and post functions return a
boolean result indicating whether to continue the traversal. Use it to write a
function ElementByID with the following signature that finds the first HTML
element with the specified id attribute. The function should stop the traversal
as soon as a match is found.

func ElementByID(doc *html.Node, id string) *html.Node
*/

package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/html"
)

func main() {

	url := flag.String("url", "", "url to parse")
	id := flag.String("id", "", "element id to find")

	
	log.SetPrefix(path.Base(os.Args[0]) + ": ")
	log.SetFlags(0)

	flag.Parse()

	if len(*url) == 0 || len(*id) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	el, err := findElementInUrl(*url, *id)

	if err != nil {
		log.Fatalf("%s not found", *id)
	}

	if el != nil {
		log.Printf("found element with id %s\n%+v",*id, *el)
	} else {
		log.Printf("not found element with id=%s", *id)
	}

}

func findElementInUrl(url, id string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := findElementInHtml(resp.Body, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func findElementInHtml(r io.Reader, id string) (*html.Node, error) {

	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return ElementByID(doc, id), nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, isHtmlNodeHasId, isHtmlNodeHasId)
}

type elementHandler func(n *html.Node, id string) bool

func forEachNode(n *html.Node, id string, pre, post elementHandler) *html.Node {

	if pre != nil && pre(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if v := forEachNode(c, id, pre, post); v != nil {
			return v
		}
	}

	if post != nil && post(n, id) {
		return n
	}
	return nil
}

func isHtmlNodeHasId(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}
