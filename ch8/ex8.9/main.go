/*

Exercise 8.9: Write a version of du that computes and periodically displays
separate totals for each of the root directories.

*/

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

func main() {

	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var walkDirWg, rootWg sync.WaitGroup
	out := make(chan string)
	outSync := make(chan struct{})
	go func() {
		for s := range out {
			fmt.Print(s)
		}
		outSync <- struct{}{}
	}()
	for _, root := range roots {
		rootWg.Add(1)
		go processRoot(&walkDirWg, &rootWg, root, out)
	}
	rootWg.Wait()
	close(out)
	<-outSync
}

func processRoot(
	walkDirWg *sync.WaitGroup,
	rootWg *sync.WaitGroup,
	root string,
	out chan<- string) {
	defer rootWg.Done()

	fileSizes := make(chan int64)

	walkDirWg.Add(1)
	go walkDir(root, walkDirWg, fileSizes)

	go func() {
		walkDirWg.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nFiles, nBytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nFiles++
			nBytes += size
		case <-tick:
			out <- printDiskUsage(root, nFiles, nBytes)
		}
	}
	out <- printDiskUsage(root, nFiles, nBytes) // final totals
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func printDiskUsage(root string, nFiles, nBytes int64) string {
	return fmt.Sprintf("%s: %d files %s\n", root, nFiles, ByteCountIEC(nBytes))
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// semaphore is a counting semaphore for limiting concurrency in dirEntries.
var semaphore = make(chan struct{}, 20)

// dirEntries returns the entries of directory dir.
func dirEntries(dir string) []os.FileInfo {
	semaphore <- struct{}{}        // acquire token
	defer func() { <-semaphore }() // release token

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}

	infos := make([]os.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "du: %v\n", err)
			continue
		}
		infos = append(infos, info)
	}

	return infos
}
