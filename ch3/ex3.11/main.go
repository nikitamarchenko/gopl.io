/*
Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers
and an optional sign.
*/

package main
import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

func validate(s string) {
	const (
		colorReset = "\033[0m"
		colorRed   = "\033[31m"
	)

	var buf bytes.Buffer
	err := false
	for _, c := range s {
		switch c {
		case '-', '+', '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			buf.WriteRune(c)
			continue
		default:
			buf.WriteString(colorRed)
			buf.WriteRune(c)
			buf.WriteString(colorReset)
			err = true
		}
	}

	if err {
		fmt.Printf("Invalid string: %s\n", buf.String())
		os.Exit(1)
	}
}

// comma inserts commas in a string.
func comma(s string) string {

	validate(s)

	var buf bytes.Buffer

	if len(s) <= 3 {
		return s
	}

	ss := strings.Split(s, ".")

	switch len(ss) {
	case 1:
		comma_subs_decimal(&buf, s)
	case 2:
		comma_subs_decimal(&buf, ss[0])
		buf.WriteRune('.')
		comma_subs_fraction(&buf, ss[1])
	default:
		fmt.Printf("to many period in string %s\n", s)
		os.Exit(1)
	}

	return buf.String()
}

func comma_subs_decimal(buf *bytes.Buffer, s string) {
	l := len(s) - 1
	for i := range s {
		buf.WriteByte(s[i])
		if i < l && (l-i)%3 == 0 && i != 0 {
			buf.WriteRune(',')
		}
	}
}

func comma_subs_fraction(buf *bytes.Buffer, s string) {
	l := len(s)
	for i := range s {
		buf.WriteByte(s[i])
		if i < l-1 && (i+1)%3 == 0 && i != 0 {
			buf.WriteRune(',')
		}
	}
}
