/*
Exercise 5.5: Implement countWordsAndImages. (See Exercise 4.9 for
word-splitting.)
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("\tUsage %s URL\n", os.Args[0])
		return
	}

	words, images, err := CountWordsAndImages(os.Args[1])

	if err != nil {
		fmt.Printf("Error %s", err)
		os.Exit(1)
	}

	fmt.Printf("words=%d images=%d\n", words, images)
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(root *html.Node) (words, images int) {

	visitor := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			images++
		} else if n.Type == html.TextNode {
			input := bufio.NewScanner(strings.NewReader(n.Data))
			input.Split(bufio.ScanWords)
			for input.Scan() {
				words++
			}
		}
	}
	visit(visitor, root)
	return
}

func visit(visitor func(*html.Node), n *html.Node) {
	if n != nil {
		visitor(n)
		visit(visitor, n.NextSibling)
		visit(visitor, n.FirstChild)
	}
}
