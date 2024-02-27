/*

Exercise 8.11: Following the approach of mirroredQuery in Section 8.4.4,
implement a variant of fetch that requests several URLs concurrently. As soon
as the first response arrives, cancel the other requests.

func mirroredQuery() string {
    responses := make(chan string, 3)
    go func() { responses <- request("asia.gopl.io") }()
    go func() { responses <- request("europe.gopl.io") }()
    go func() { responses <- request("americas.gopl.io") }()
    return <-responses // return the quickest response
}

*/

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	response := make(chan string, len(os.Args)-1)
	abort := make(chan struct{})
	var wg sync.WaitGroup
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go fetch(url, &wg, response, abort)
	}
	fmt.Println(<-response)
	close(abort)
	wg.Wait()
}

func fetch(url string, wg *sync.WaitGroup, response chan<- string, abort chan struct{}) {
	defer wg.Done()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Cancel = abort
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return
	}
	_ = b
	//fmt.Printf("%s", b)
	response <- url
}
