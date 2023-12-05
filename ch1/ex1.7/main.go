/*
Exercis e 1.7: The function call io.Copy(dst, src) reads from src and w rites to dst.
Use it instead o f ioutil.ReadAll to copy the response body to os.Stdout without
requiring a buffer large enough to hold t he entire stream. Be sure to chek the
error result of io.Copy
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		_, err =io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}
}

//!-
