/*

Exercise 7.17: Extend xmlselect so that elements may be selected not just by
name, but by their attributes too, in the manner of CSS, so that, for instance,
an element like <div id="page" class="wide"> could be selected by a matching
id or class as well as its name.
*/

package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Attrs map[string]string

type Token struct {
	name  string
	attrs Attrs
}

func (s Token) String() string {

	if len(s.attrs) == 0 {
		return s.name
	}

	attrs := make([]string, 0, len(s.attrs))
	for k, v := range s.attrs {
		attrs = append(attrs, fmt.Sprintf("%s=%s", k, v))
	}

	return fmt.Sprintf("%s[%s]", s.name, strings.Join(attrs, ", "))
}

func NewToken(el xml.StartElement) (t Token) {
	t.name = el.Name.Local
	t.attrs = make(map[string]string)
	for _, a := range el.Attr {
		t.attrs[a.Name.Local] = a.Value
	}
	return
}

func (t Token) Check(s string) (r bool, err error) {

	// no selectors just name
	if !strings.Contains(s, "[") {
		return t.name == s, nil
	}

	// selector
	fb := strings.Index(s, "[")
	lb := strings.Index(s, "]")

	if t.name != s[:fb] {
		return false, nil
	}

	if lb == -1 || fb > lb {
		return false, fmt.Errorf("error: invalid selector format %s", s)
	}

	// selector by attribute without value
	sel := s[fb+1 : lb]
	if len(sel) == 0 {
		return true, nil
	}

	if strings.Contains(sel, "=") {
		subs := strings.Split(sel, "=")
		if v, ok := t.attrs[subs[0]]; ok && v == subs[1] {
			return true, nil
		}
	} else {
		if _, ok := t.attrs[sel]; ok {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []Token // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:

			stack = append(stack, NewToken(tok)) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {

				b := make([]string, 0, len(stack))
				for _, v := range stack {
					b = append(b, v.String())
				}

				fmt.Printf("%s: %s\n", strings.Join(b, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []Token, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		r, err := x[0].Check(y[0])
		if err != nil {
			fmt.Printf("error %s", err)
			return false
		}
		if r {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

//!-
