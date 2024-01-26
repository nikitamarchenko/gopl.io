// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

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

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	if l := x.Len(); l != 4 {
		t.Errorf("len 4 != %d \n", l)
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	x.Remove(10000)

	if s := x.String(); s != "{1 9 42 144}" {
		t.Errorf("broken delete")
	}

	x.Remove(1)

	if x.Has(1) {
		t.Errorf("broken delete 1")
	}

	x.Remove(42)

	if x.Has(42) {
		t.Errorf("broken delete 42")
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	x.Clear()

	if x.Len() != 0 {
		t.Error("error: clear func")
	}

}
func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	r := x.Copy()

	if &x == r {
		t.Error("error: we copied entire object")
	}

	if x.String() != r.String() {
		t.Error("error: object is different")
	}
}

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
   
	y.Add(2)
	y.Add(3)
	y.Add(4)
   
	fmt.Println("x:", x.String()) // "{1 2 3}"
	fmt.Println("y:", y.String()) // "{2 3 4}"
   
	x.IntersectWith(&y)
	if r := x.String(); r != "{2 3}" {
		t.Errorf("error: intersect %s\n", r)
	}
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
   
	y.Add(2)
	y.Add(3)
	y.Add(4)
   
	x.DifferenceWith(&y)

	if r := x.String(); r != "{1}" {
		t.Errorf("error: difference with %s\n", r)
	}
}


func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(2)
	x.Add(3)
   
	y.Add(2)
	y.Add(3)
	y.Add(4)
   
	x.SymmetricDifference(&y)

	if r := x.String(); r != "{1 4}" {
		t.Errorf("error: symmetric difference %s\n", r)
	}
}

