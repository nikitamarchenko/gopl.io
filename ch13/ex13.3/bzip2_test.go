// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bzip

import (
	"bufio"
	"bytes"
	"compress/bzip2" // reader
	"fmt"
	"io"
	"sync"
	"testing"
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := NewWriter(&compressed)

	// Write a repetitive message in a million pieces,
	// compressing one copy but not the other.
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Check the size of the compressed stream.
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}

	// Decompress and compare with original.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decompression yielded a different message")
	}
}

func TestBzip2MultiWrite(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := NewWriter(&compressed)

	data := make(map[string]bool)
	ch := make(chan string, 100)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for s := range ch {
				w.Write([]byte(s))
			}
			wg.Done()
		}()
	}

	for i := 0; i < 1000000; i++ {
		k := fmt.Sprintf("{%d}", i)
		s := fmt.Sprintf("%s\n", k)
		uncompressed.WriteString(s)
		ch <- s
		data[k] = true
	}
	
	close(ch)
	wg.Wait()
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	// Decompress and compare with original.
	var decompressed bytes.Buffer
	io.Copy(&decompressed, bzip2.NewReader(&compressed))
	s := bufio.NewScanner(bytes.NewReader(decompressed.Bytes()))
	for s.Scan() {
		text := s.Text()
		if !data[text] {
			t.Fatal("invalid text compressed")
		}
		delete(data, text)
	}
	if len(data) > 0 {
		t.Fatalf("not all data compressed missed %d", len(data))
	}
}
