// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestString(t *testing.T) {
	tr := add(nil, 1)
	tr = add(tr, 2)
	tr = add(tr, 3)
	tr = add(tr, 4)


	if fmt.Sprint(tr) != "[1(2(3(4)))]" {
		t.Errorf("invalid string representation: %s", tr)
	}

	tr = add(nil, 3)
	tr = add(tr, 2)
	tr = add(tr, 4)

	if fmt.Sprint(tr) != "[3(2)(4)]" {
		t.Errorf("invalid string representation: %s", tr)
	}
}
