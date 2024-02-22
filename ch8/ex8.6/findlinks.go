/*

Exercise 8.6: Add depth-limiting to the concurrent crawler. That is, if the
user sets -depth=3, then only URLs reachable by at most three links will be
fetched.

*/

package main

import (
	"flag"
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"sync/atomic"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

type Link struct {
	url   string
	depth int
}

func makeLinks(depth int, links []string) []Link {
	wl := make([]Link, len(links))
	for i, l := range links {
		wl[i] = Link{l, depth}
	}
	return wl
}

func main() {
	depth := flag.Int("depth", 1, "if the user sets -depth=3, then only URLs reachable by at most three links will be fetched.")
	flag.Parse()
	worklist := make(chan []Link)  // lists of URLs, may have duplicates
	unseenLinks := make(chan Link) // de-duplicated URLs

	// Add command-line arguments to worklist.
	works := atomic.Int32{}
	wl := makeLinks(0, flag.Args())
	if len(wl) > 0 {
		works.Add(1)
		go func() { worklist <- wl }()
	} else {
		return
	}

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := makeLinks(link.depth+1, crawl(link.url))
				if len(foundLinks) > 0 {
					works.Add(1)
					go func(foundLinks []Link) { worklist <- foundLinks }(foundLinks)
				}
				works.Add(-1)
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		works.Add(-1)
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				if link.depth < *depth {
					works.Add(1)
					unseenLinks <- link
				}
			}
		}
		if works.Load() == 0 {
			close(unseenLinks)
			close(worklist)
		}
	}
}
