package main

import (
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {

	r := strings.NewReader("0123456789;01234567890")
	lr := LimitReader(r, 11)
	b := make([]byte, 11)
	if n, _ := lr.Read(b); n != 11 {
		t.Errorf("invalid read len %d", n)
	}
	if string(b) != "0123456789;" {
		t.Errorf("invalid text %s", string(b))
	}
	if n, _ := lr.Read(b); n != 11 {
		t.Errorf("invalid second read %d", n)
	}
	if n, err := lr.Read(b); n != 0 || err != io.EOF {
		t.Errorf("invalid read eof %d, %s", n, err)
	}
}
