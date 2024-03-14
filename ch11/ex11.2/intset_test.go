/*
Exercise 11.2: Write a set of tests for IntSet (§6.5) that checks that its
behavior after each operation is equivalent to a set based on built-in maps.
Save your implementation for benchmarking in Exercise 11.7.
*/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func buildSets(bits []int) (IntSet, IntSetMap) {
	var x IntSet
	var y IntSetMap

	for _, v := range bits {
		x.Add(v)
		y.Add(v)
	}

	if x.String() != y.String() {
		panic("Not equal")
	}
	return x, y
}

func TestIntSet_Has(t *testing.T) {
	tests := []struct {
		name  string
		args  []int
		check []int
	}{
		{"{1, 144, 9, 42}", []int{1, 144, 9, 42}, []int{1, 144, 9, 42, 10, 15, 120, 150}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y := buildSets(tt.args)
			for _, v := range tt.check {
				if x.Has(v) != y.Has(v) {
					t.Errorf("Has(%d) got = %v, want %v", v, x.Has(v), y.Has(v))
				}
			}
		})
	}
}

func TestIntSet_UnionWith(t *testing.T) {

	type args struct {
		l, r []int
	}

	tests := []struct {
		name string
		args args
	}{
		{"{1, 144, 9, 42} | {15, 32, 100}", args{l: []int{1, 144, 9, 42}, r: []int{15, 32, 100}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lx, ly := buildSets(tt.args.l)
			rx, ry := buildSets(tt.args.r)

			lx.UnionWith(&rx)
			ly.UnionWith(ry)

			if lx.String() != ly.String() {
				t.Errorf("UnionWith %v | %v, got %s want %s", tt.args.l, tt.args.r, &lx, &ly)
			}
		})
	}
}
