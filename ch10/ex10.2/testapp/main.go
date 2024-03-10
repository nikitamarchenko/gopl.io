package main

import (
	"fmt"
	"os"

	"gopl.io/ch10/ex10.2/archive"
	_ "gopl.io/ch10/ex10.2/archive/zip"
	_ "gopl.io/ch10/ex10.2/archive/tar"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error: path not specified")
		os.Exit(1)
	}

	r, err := archive.NewReader(os.Args[1])
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	defer r.Close()

	for r.Next() {
		f := r.Get()
		fmt.Println(f)
	}

	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
