/*

Exercise 12.4: Modify encode to pretty-print the S-expression in the style shown
above.

((Title "Dr. Strangelove")
 (Subtitle "How I Learned to Stop Worrying and Love the Bomb")
 (Year 1964)
 (Actor (("Grp. Capt. Lionel Mandrake" "Peter Sellers")
         ("Pres. Merkin Muffley" "Peter Sellers")
         ("Gen. Buck Turgidson" "George C. Scott")
         ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
         ("Maj. T.J. \"King\" Kong" "Slim Pickens")
         ("Dr. Strangelove" "Peter Sellers")))
 (Oscars ("Best Actor (Nomin.)"
          "Best Adapted Screenplay (Nomin.)"
          "Best Director (Nomin.)"
          "Best Picture (Nomin.)"))
 (Sequel nil))

*/
// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

// !+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	pw := prettyWriter{buf: &buf}
	if err := encode(&pw, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type prettyWriter struct {
	buf *bytes.Buffer
	s   []int
	sc  int
}

func (pw *prettyWriter) Write(p []byte) (n int, err error) {
	return pw.buf.Write(p)
}

func (pw *prettyWriter) WriteByte(c byte) error {
	_, err := pw.Write([]byte{c})
	return err
}

func (pw *prettyWriter) pushSpace(i int) {
	pw.s = append(pw.s, i)
	pw.sc += i
}

func (pw *prettyWriter) popSpace() {
	pw.sc -= pw.s[len(pw.s)-1]
	pw.s = pw.s[:len(pw.s)-1]
}

func (pw *prettyWriter) writeSpace() {
	for i := 0; i < pw.sc; i++ {
		pw.buf.WriteByte(' ')
	}
}

//io.Writer

//!-Marshal

// encode writes to buf an S-expression representation of v.
// !+encode
func encode(buf *prettyWriter, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprint(buf, "nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		buf.pushSpace(1)
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.writeSpace()
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
			if i < v.Len()-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')
		buf.popSpace()

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		buf.pushSpace(1)
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.writeSpace()
			}
			s := fmt.Sprintf("(%s ", v.Type().Field(i).Name)
			buf.pushSpace(len(s))
			fmt.Fprint(buf, s)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.popSpace()
			buf.WriteByte(')')
			if i < v.NumField()-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')
		buf.popSpace()

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		buf.pushSpace(1)
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.writeSpace()
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
			if i < len(v.MapKeys())-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')
		buf.popSpace()

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

//!-encode
