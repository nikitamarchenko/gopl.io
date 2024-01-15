/*
Exercise 5.13: Modify crawl to make local copies of the pages it finds,
creating directories as necessary. Don’t make copies of pages that come from a
different domain. For example, if the original page comes from golang.org,
save all files from there, but exclude ones from vimeo.com.
*/

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"errors"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"net/http"
	"net/url"

	"gopl.io/ch5/links"
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

//!-breadthFirst

// !+crawl
func crawl(u string) []string {
	log.Println(u)
	startingUrl, err := url.Parse(u)
	if err != nil {
		log.Print(err)
	}
	startingUrl.Fragment = ""

	if startingUrl.Scheme != "https" && startingUrl.Scheme != "http" {
		log.Printf("skip bc scheme %s not valid", startingUrl.Scheme)
		return []string{}
	}
	req, _ := http.NewRequest(http.MethodGet, startingUrl.String(), nil)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		return []string{}
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("request url %s not ok status %s", u, resp.Status)
	}

	var valid bool
	for name, values := range resp.Header {
		if name == "Content-Type" {
			for _, value := range values {
				if strings.Contains(value, "text/") {
					valid = true
				}
			}
		}
	}

	if !valid {
		log.Printf("url %s is not valid type", u)
		return []string{}
	}

	pathChunks := strings.Split(startingUrl.EscapedPath(), "/")

	var isFile bool
	if len(pathChunks) > 0 {
		// that't really stupid but we can't know for sure is it file or not
		if strings.Contains(pathChunks[len(pathChunks)-1], ".") {
			isFile = true
		}
	}

	rootDir := path.Join(path.Dir("."), startingUrl.Hostname())
	err = os.MkdirAll(rootDir, 0775)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	var fileDir, filename string

	if isFile {
		filename = pathChunks[len(pathChunks)-1]
		fileDir = path.Join(pathChunks[:len(pathChunks)-1]...)
	} else {
		filename = "index.html"
		fileDir = path.Join(pathChunks...)
	}

	fileDir = path.Join(rootDir, fileDir)

	log.Println("dir:", fileDir, "filename", filename)

	err = os.MkdirAll(fileDir, 0775)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	var download bool
	filepath := path.Join(fileDir, filename)
	if _, err := os.Stat(filepath); err == nil {
		log.Println("already downloaded")
	} else if errors.Is(err, os.ErrNotExist) {
		download = true
	} else {
		log.Printf("error on file %s\n", err)
	}

	// good way to to this in links.Extract but we can only modify crawl method
	if download {
		func() {
			resp, err = http.Get(u)
			if err != nil {
				log.Printf("download %s error: %s\n", u, err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				
				log.Printf("download %s error status: %s\n", u, resp.Status)
				return
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("download %s error: %s\n", u, err)
				return
			}
			err = os.WriteFile(filepath, body, 0664)
			if err != nil {
				log.Printf("save %s error: %s\n", u, err)
				return
			}
		}()
	}

	list, err := links.Extract(u)
	if err != nil {
		log.Print(err)
		return []string{}
	}

	hostname := startingUrl.Hostname()
	for i, l := range list {
		if strings.HasPrefix(l, "/") {
			r, err := url.JoinPath(hostname, l)
			if err != nil {
				log.Printf("error normalize url %s with hostname %s",
					l, hostname)
			} else {
				list[i] = r
			}
		}
	}

	result := make([]string, 0, len(list))
	for _, u := range list {
		newUrl, err := url.Parse(u)
		if err != nil {
			log.Print(err)
			continue
		}
		if newUrl.Hostname() == "" {
			log.Println("skip", newUrl)
		}
		if strings.HasPrefix(newUrl.Hostname(), hostname) {
			result = append(result, u)
		}
	}

	return result
}

//!-crawl

// !+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main
