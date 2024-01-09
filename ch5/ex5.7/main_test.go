package main

import (
	"strings"
	"testing"
)

func TestProcessHtml(t *testing.T) {

	var tests = []struct {
		name    string
		html    string
		want    []string
		inverse bool
	}{
		{"div with args 1",
			`<div arg1='val1'></div>`,
			[]string{` arg1="val1"`},
			false},
		{"div with args 2",
			`<div arg1='val1' arg2='val2'></div>`,
			[]string{` arg1="val1"`, ` arg2="val2"`},
			false},
		{"div no children",
			`<div></div>`,
			[]string{`<div/>`},
			false},
		{"div no children",
			`<div></div>`,
			[]string{`</div>`},
			true},
		{"p with text",
			`<p>text</p>`,
			[]string{`text`},
			false},
		{"skip empty text nodes",
			"<div><p>text</p>\n</div>",
			[]string{"\n      \n"},
			true},
		{"comment",
			"<!--This is a comment-->",
			[]string{"<!--This is a comment-->"},
			false},
	}

	for _, test := range tests {
		got, err := processHtml(strings.NewReader(test.html))
		if err != nil {
			t.Errorf("%s failed: %v", test.name, err)
			continue
		}
		for _, want := range test.want {
			if !strings.Contains(got, want) && !test.inverse {
				t.Errorf("%s failed: must contain %s\n%s\n",
					test.name, want, got)
			}
			if strings.Contains(got, want) && test.inverse {
				t.Errorf("%s failed: must not contain %s\n%s\n",
					test.name, want, got)
			}
		}
	}
}
