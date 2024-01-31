/*

Exercise 7.2: Write a function CountingWriter with the signature below that,
given an io.Writer, returns a new Writer that wraps the original, and a
pointer to an int64 variable that at any moment contains the number of bytes
written to the new Writer.


func CountingWriter(w io.Writer) (io.Writer, *int64)

*/

package main

import (
	"fmt"
	"io"
	"os"
)

type countingWriterWrapper struct {
	counter *int64
	wrap io.Writer 
}


func (cwd *countingWriterWrapper) Write(p []byte) (n int, err error) {
	n, err = cwd.wrap.Write(p)
	*cwd.counter += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	r := countingWriterWrapper{wrap: w, counter: new(int64)}
	return &r, r.counter
}


func main () {
	w :=  io.Writer(os.Stdout)
	w2, counter := CountingWriter(w)
	fmt.Fprint(w2, "0123456789")
	fmt.Println("\ncounter", *counter)
	fmt.Fprint(w2, "0123456789")
	fmt.Println("\ncounter", *counter)
}
