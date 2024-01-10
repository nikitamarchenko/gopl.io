package main

import (
	"strings"
	"testing"
)

func TestProcessHtml(t *testing.T) {

	var tests = []struct {
		name  string
		html  string
		id    string
		found bool
	}{
		{"div id=val1",
			`<div id='val1'></div>`,
			`val1`,
			true},

		{"not found",
			`<div id='val1'></div>`,
			`val2`,
			false},
	}

	for _, test := range tests {
		got, err := findElementInHtml(strings.NewReader(test.html), test.id)
		if err != nil {
			t.Errorf("%s failed: %v", test.name, err)
			continue
		}

		if test.found && got == nil {
			t.Errorf("%s failed: %s not found", test.name, test.id)
		}
	}
}
