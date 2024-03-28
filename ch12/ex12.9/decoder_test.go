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
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {

	data := `
((Title "Dr. Strangelove")
         (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Year 1964)
         (Actor
          (("Maj. T.J. \"King\" Kong" "Slim Pickens")
           ("Dr. Strangelove" "Peter Sellers")
           ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
           ("Pres. Merkin Muffley" "Peter Sellers")
           ("Gen. Buck Turgidson" "George C. Scott")
           ("Brig. Gen. Jack D. Ripper" "Sterling Hayden")))
         (Oscars
          ("Best Actor (Nomin.)" "Best Adapted Screenplay (Nomin.)"
           "Best Director (Nomin.)" "Best Picture (Nomin.)")) (Sequel nil))
`
	dec := NewDecoder(strings.NewReader(data))

	want := []interface{}{
		StartList{},
		StartList{},
		Symbol("Title"),
		String("Dr. Strangelove"),
		EndList{},
		StartList{},
		Symbol("Subtitle"),
		String("How I Learned to Stop Worrying and Love the Bomb"),
		EndList{},
		StartList{},
		Symbol("Year"),
		Int(1964),
		EndList{},
		StartList{},
		Symbol("Actor"),
		StartList{},
		StartList{},
		String("Maj. T.J. \"King\" Kong"),
		String("Slim Pickens"),
		EndList{},
		StartList{},
		String("Dr. Strangelove"),
		String("Peter Sellers"),
		EndList{},
		StartList{},
		String("Grp. Capt. Lionel Mandrake"),
		String("Peter Sellers"),
		EndList{},
		StartList{},
		String("Pres. Merkin Muffley"),
		String("Peter Sellers"),
		EndList{},
		StartList{},
		String("Gen. Buck Turgidson"),
		String("George C. Scott"),
		EndList{},
		StartList{},
		String("Brig. Gen. Jack D. Ripper"),
		String("Sterling Hayden"),
		EndList{},
		EndList{},
		EndList{},
		StartList{},
		Symbol("Oscars"),
		StartList{},
		String("Best Actor (Nomin.)"),
		String("Best Adapted Screenplay (Nomin.)"),
		String("Best Director (Nomin.)"),
		String("Best Picture (Nomin.)"),
		EndList{},
		EndList{},
		StartList{},
		Symbol("Sequel"),
		Symbol("nil"),
		EndList{},
		EndList{},
	}

	for i, v := range want {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Errorf("%v\n", err)
		}
		switch tok := tok.(type) {
		case Symbol:
			t.Logf("%T %v", tok, tok)
		case String:
			t.Logf("%T %v", tok, tok)
		case Int:
			t.Logf("%T %v", tok, tok)
		case StartList:
			t.Logf("%T %v", tok, tok)
		case EndList:
			t.Logf("%T %v", tok, tok)
		}
		if !reflect.DeepEqual(tok, v) {
			t.Errorf("i = %d got %v want %v", i, tok, v)
		}
	}
}
