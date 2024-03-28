/*

Exercise 12.9: Write a token-based API for decoding S-expressions, following
the style of xml.Decoder (§7.14). You will need five types of tokens: Symbol,
String, Int, StartList, and EndList.

func main() {
    dec := xml.NewDecoder(os.Stdin)
    var stack []string // stack of element names
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
            stack = append(stack, tok.Name.Local) // push
        case xml.EndElement:
            stack = stack[:len(stack)-1] // pop
        case xml.CharData:
            if containsAll(stack, os.Args[1:]) {
                fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
            }
        }
    }
}
*/

package sexpr

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type Decoder struct {
	s    *scanner.Scanner
	sErr string
}

func NewDecoder(i io.Reader) *Decoder {
	d := &Decoder{}
	s := &scanner.Scanner{Mode: scanner.GoTokens}
	s.Error = func(s *scanner.Scanner, msg string) {
		d.sErr = msg
	}
	s.Init(i)
	d.s = s
	return d
}

type Token interface{}
type Symbol string
type String string
type Int int
type StartList struct{}
type EndList struct{}

func (StartList) String() string {
	return "("
}

func (EndList) String() string {
	return ")"
}

func (d *Decoder) tt() string {
	return d.s.TokenText()
}

func (d *Decoder) Token() (Token, error) {
	if d.sErr != "" {
		return nil, fmt.Errorf("scanner err: %s", d.sErr)
	}

	t := d.s.Scan()
	switch t {
	case scanner.Ident:
		return Symbol(d.tt()), nil
	case scanner.String:
		s, err := strconv.Unquote(d.tt())
		if err != nil {
			return nil, fmt.Errorf("decoder: %v", err)
		}
		return String(s), nil
	case scanner.Int:
		i, err := strconv.Atoi(d.tt())
		if err != nil {
			return nil, fmt.Errorf("decoder can't parse int %s", d.tt())
		}
		return Int(i), nil
	case '(':
		return StartList{}, nil
	case ')':
		return EndList{}, nil
	case scanner.EOF:
		return nil, io.EOF
	}

	return nil, fmt.Errorf("unexpected token %q", d.tt())
}
