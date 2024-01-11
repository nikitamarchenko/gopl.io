/*
Exercise 5.9: Write a function expand(s string, f func(string) string) string
that replaces each substring “$foo” within s by the text returned by f("foo").
*/

package main

import (
	"bytes"
	"strings"
	"unicode"
)

func trivialExpand(s string, f func(string) string) string {
	words := strings.Split(s, " ")

	for i, w := range words {
		if after, found := strings.CutPrefix(w, "$"); found {
			words[i] = f(after)
		}
	}

	return strings.Join(words, " ")
}

func expand(s string, f func(string) string) string {

	const (
		scanning byte = iota
		collect
	)

	var buf, val bytes.Buffer
	state := scanning
	for _, r := range s {
		switch state {
		case scanning:
			if r == '$' {
				state = collect
			} else {
				buf.WriteRune(r)
			}
		case collect:
			if r == '$' {
				buf.WriteString(f(val.String()))
				val.Reset()
			} else if unicode.IsSpace(r) {
				buf.WriteString(f(val.String()))
				buf.WriteRune(r)
				val.Reset()
				state = scanning
			} else {
				val.WriteRune(r)
			}
		}
	}

	if val.Len() > 0 {
		buf.WriteString(f(val.String()))
	}

	return buf.String()
}
