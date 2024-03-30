/*

Exercise 12.11: Write the corresponding Pack function. Given a struct value,
Pack should return a URL incorporating the parameter values from the struct.

*/

package pack

import "testing"

func TestPack(t *testing.T) {

	type P struct {
		Int     int
		Uint    uint
		Float64 float64
		Bool    bool
		String  string
		TaggedInt int `http:"tagged"`
	}

	data := P{-42, 42, 42.42, true, "my/cool+blog&about,stuff", 42*42}

	got, err := Pack(&data)

	if err != nil {
		t.Fatalf("err %v", err)
	}

	want := `int=-42&uint=42&float64=42.42&bool=true&string="my%2Fcool%2Bblog%26about%2Cstuff"&tagged=1764`
	if got != want {
		t.Fatalf("got %s want %s", got, want)
	} 
}
