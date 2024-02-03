/*
Exercise 7.5: The LimitReader function in the io package accepts an io.Reader
r and a number of bytes n, and returns another Reader that reads from r but
reports an end-of-file condition after n bytes. Implement it.

func LimitReader(r io.Reader, n int64) io.Reader

*/

package main

import (
	"io"
)

type limitReaderData struct {
	r *io.Reader
	n int64
}

func (l *limitReaderData) Read(p []byte) (n int, err error) {
	return (*l.r).Read(p[:l.n])
}


func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReaderData{r:&r, n:n}
}
