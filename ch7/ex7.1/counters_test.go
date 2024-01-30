package counters

import (
	"fmt"
	"testing"
)


func TestLineCounter(t *testing.T) {

	var lc LineCounter

	if r, _ := lc.Write([]byte("1\n2\n3\n")); r != 3 {
		t.Error("invalid count")
	}

	if lc != 3 {
		t.Error("invalid count")
	}

	if fmt.Sprint(&lc) != "count = 3" {
		t.Errorf("invalid count")
	}
}

func TestWordCounter(t *testing.T) {

	var wc WordCounter

	if r, _ := wc.Write([]byte("1 2 3")); r != 3 {
		t.Errorf("invalid count %d", r)
	}

	if wc != 3 {
		t.Error("invalid count")
	}

	if fmt.Sprint(&wc) != "count = 3" {
		t.Errorf("invalid count")
	}
}

