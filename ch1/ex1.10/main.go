/*

Exercise 1.10: Find a web site that produces a large amount of data. Investigate chanhing
by running fetchall twice in succession to see whether the reported time changes much.
Do you get the same coontent each time? Modify fetchall to print its out put to a
file so it can be examined.

*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
    start := time.Now()
    ch := make(chan string)

    newpath := filepath.Join(".", "out")

    err := os.MkdirAll(newpath, os.ModePerm)

    if err != nil {
        fmt.Printf("Error while create dir out %s", err)
        return
    }

    for _, url := range os.Args[1:] {
        go fetch(url, ch) // start a goroutine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch) // receive from channel ch
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)

    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }
    defer resp.Body.Close() // don't leak resources

    err, dst_filename := FilenameFromUrl(url, ch)

    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }

    f, err := os.Create(dst_filename)
    defer f.Close()

    if err != nil {
        ch <- fmt.Sprint(err) // send to channel ch
        return
    }

    w := bufio.NewWriter(f)
    nbytes, err := io.Copy(w, resp.Body)
    
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
    w.Flush()
}

func FilenameFromUrl(urlstr string, ch chan<- string) (error, string) {
    u, err := url.Parse(urlstr)
    if err != nil {
        ch <- fmt.Sprintf("Error due to parsing url: %s", urlstr)
        return err, ""
    }
    x, _ := url.QueryUnescape(u.EscapedPath())
    url_to_path := u.Hostname() + strings.ReplaceAll(filepath.Clean(x), "/", "_")
    f := filepath.Join(".", "out", url_to_path)
    return nil, f
}

//!-
