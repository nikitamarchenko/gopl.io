/*

Exercise 10.1: Extend the jpeg program so that it converts any supported input
format to any output format, using image.Decode to detect the input format and
a flag to select the output format.

*/

package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	to := flag.String("t", "jpeg", "jpeg, png, gif")
	flag.Parse()
	switch *to {
	case "jpeg", "gif", "png":
	default:
		fmt.Println("error: invalid format")
		flag.PrintDefaults()
		return
	}

	if err := convert(*to, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func convert(to string, in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)

	switch to {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 256})
	case "png":
		return png.Encode(out, img)
	}
	return fmt.Errorf("convert format")
}
