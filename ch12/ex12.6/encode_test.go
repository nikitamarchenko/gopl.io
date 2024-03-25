package sexpr

import (
	"testing"
)

func Test(t *testing.T) {
	type S struct {
		DefaultTest string
		F           int
	}
	s := S{
		DefaultTest: "",
		F:           1,
	}

	// Encode it
	data, err := Marshal(s)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	want := "((F 1))"
	if string(data) != want {
		t.Fatalf("want = %s, got = %s", want, data)
	}

}
