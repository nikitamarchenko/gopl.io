/*

Exercise 12.7: Create a streaming API for the S-expression decoder, following
the style of json.Decoder (§4.5).

*/

package sexpr

import (
	"io"
	"strings"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {

	const sStream = `
	42
	"text"
	(1 2 3 4 5)
	((F 1))
`
	type S struct {
		F int
	}

	dec := NewDecoder(strings.NewReader(sStream))
	var i int
	if err := dec.Decode(&i); err == io.EOF {
		t.Fatal(err)
	}
	t.Logf("%T: %v\n", i, i)

	var s string
	if err := dec.Decode(&s); err == io.EOF {
		t.Fatal(err)
	}
	t.Logf("%T: %v\n", s, s)

	var sl []int
	if err := dec.Decode(&sl); err == io.EOF {
		t.Fatal(err)
	}
	t.Logf("%T: %v\n", sl, sl)

	var st S
	if err := dec.Decode(&st); err == io.EOF {
		t.Fatal(err)
	}
	t.Logf("%T: %v\n", st, st)

	var eof interface{}
	if err := dec.Decode(&eof); err != io.EOF {
		t.Fatal("not eof")
	}
}
