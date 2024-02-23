/*

Exercise 8.7: Write a concurrent program that creates a local mirror of a web
site, fetching each reachable page and writing it to a directory on the local
disk. Only pages within the original domain (for instance, golang.org) should
be fetched. URLs within mirrored pages should be altered as needed so that
they refer to the mirrored page, not the original.

*/

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"net/http"
	"net/url"

	"sync/atomic"

	"gopl.io/ch5/links"
)

const ()

func breadthFirst(f func(item string) []string, url string) {
	jobs := make(chan string)
	results := make(chan []string)
	var jobsCount atomic.Int32
	for concurrent := runtime.NumCPU() * 2; concurrent > 0; concurrent-- {
		go func() {
			for link := range jobs {
				go func (ss []string) {results <- ss}(f(link))
				jobsCount.Add(-1)
			}
		}()
	}
	jobsCount.Add(1)
	go func() { jobs <- url }()
	seen := make(map[string]bool)
	for links := range results {
		for _, l := range links {
			if !seen[l] {
				seen[l] = true
				jobsCount.Add(1)
				jobs <- l
			}
		}
		if jobsCount.Load() == 0 {
			break
		}
	}
}

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
		return []string{}
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

func main() {
	if len(os.Args) == 2 {
		breadthFirst(crawl, os.Args[1])
	} else {
		fmt.Printf("usage %s <url>", os.Args[0])
	}
}
