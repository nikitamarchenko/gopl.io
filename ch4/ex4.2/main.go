/*
Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by
default but supports a command-line flag to print the SHA384 or SHA512 hash instead.
*/

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"flag"
	"fmt"
	"hash"
	"strings"
)

type shaType int

const (
	SHA256 shaType = iota
	SHA384
	SHA512
)

var shaFlag shaType

func init() {

	flag.Func("type", "SHA256 SHA384 SHA512", func(s string) error {

		s = strings.ToUpper(s)

		switch s {
		case "SHA256":
			shaFlag = SHA256
		case "SHA384":
			shaFlag = SHA384
		case "SHA512":
			shaFlag = SHA512
		default:
			return errors.New("could not parse type of sha")
		}

		return nil
	})

}

func main() {

	flag.Parse()

	hf := [...]hash.Hash{sha256.New(), sha512.New384(), sha512.New()}

	sum := hf[shaFlag].Sum([]byte(flag.Arg(0)))

	fmt.Printf("%x\n", sum)
}
