package main

import "testing"

func TestExpand(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		f      func(string) string
		output string
	}{
		{name: "empty"},
		{name: "basic",
			input: "$foo",
			f: func(s string) string {
				return "baz"
			},
			output: "baz",
		},
		{name: "basic",
			input: "$foo $test",
			f: func(s string) string {
				return s + "!"
			},
			output: "foo! test!",
		},
		{name: "intermediate",
			input: " _  __$foo$test  _",
			f: func(s string) string {
				return s + "!"
			},
			output: " _  __foo!test!  _",
		},
		{name: "advanced",
			input: "  start$你好$世界 end ",
			f: func(s string) string {
				return "(" + s + ")"
			},
			output: "  start(你好)(世界) end ",
		},
	}

	for _, test := range tests {
		if r := expand(test.input, test.f); r != test.output {
			t.Errorf("error in %s:\n\tinput  : [%s]\n\toutput : [%s]\n\tresult = [%s]",
				test.name, test.input, test.output, r)
		} else {
			t.Logf("%s: pass", test.name)
		}
	}
}


func TestTrivialExpand(t *testing.T) {

	tests := []struct {
		name   string
		input  string
		f      func(string) string
		output string
	}{
		{name: "empty"},
		{name: "basic",
			input: "$foo",
			f: func(s string) string {
				return "baz"
			},
			output: "baz",
		},
		{name: "basic",
			input: "$foo $test",
			f: func(s string) string {
				return s + "!"
			},
			output: "foo! test!",
		},
	}

	for _, test := range tests {
		if r := trivialExpand(test.input, test.f); r != test.output {
			t.Errorf("error in %s:\n\tinput  : [%s]\n\toutput : [%s]\n\tresult = [%s]",
				test.name, test.input, test.output, r)
		} else {
			t.Logf("%s: pass", test.name)
		}
	}
}