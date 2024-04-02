/*

Exercise 13.1: Define a deep comparison function that considers numbers (of any
type) equal if they differ by less than one part in a billion.

*/

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 359.

// Package equal provides a deep equivalence relation for arbitrary values.
package equal

import (
	"testing"
)

func TestEqual(t *testing.T) {

	fnum1 := 3.14159265359
	fnum2 := 3.14159265359 + 1e-10 // differ by one part in a billion
	fnum3 := 2.71828182846

	unum1 := uint64(1000000000)
	unum2 := uint64(1000000001) // differ by one part in a billion
	unum3 := uint64(2000000000)

	num1 := int64(1000000000)
	num2 := int64(1000000001) // differ by one part in a billion
	num3 := int64(2000000000)

	cnum1 := complex(1.0, 2.0)
	cnum2 := complex(1.0+1e-10, 2.0+1e-10) // differ by one part in a billion
	cnum3 := complex(3.0, 4.0)

	type args struct {
		x interface{}
		y interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Float true", args{fnum1, fnum2}, true},
		{"Float false", args{fnum1, fnum3}, false},
		{"UInt true", args{unum1, unum2}, true},
		{"UInt false", args{unum1, unum3}, false},
		{"Int true", args{num1, num2}, true},
		{"Int false", args{num1, num3}, false},
		{"Complex true", args{cnum1, cnum2}, true},
		{"Complex false", args{cnum1, cnum3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bool
			if got = Equal(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Equal(%v, %v) = %v, want %v", tt.args.x, tt.args.y, got, tt.want)
			}
			t.Logf("Equal(%v, %v) = %v, want %v", tt.args.x, tt.args.y, got, tt.want)
		})
	}
}
