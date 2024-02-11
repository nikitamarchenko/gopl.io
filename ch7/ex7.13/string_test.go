package eval

import (
	"testing"
)

func TestString(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{input: "+A"},
		{input: "sqrt(A / pi)"},
		{input: "pow(x, 3) + pow(y, 3)"},
		{input: "5 / 9 * (F - 32)"},
	}

	for i, test := range tests {
		expr, err := Parse(test.input)
		if err != nil {
			t.Errorf("parse error %s", err)
			continue
		}
		tests[i].want = expr.String()
	}

	for _, test := range tests {
		expr, err := Parse(test.input)
		if err != nil {
			t.Errorf("parse error %s", err)
			continue
		}
		got := expr.String()
		if got != test.want {
			t.Errorf("%s => %s, want %s",
				test.input, got, test.want)
		}
	}
}
