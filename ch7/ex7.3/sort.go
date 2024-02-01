/*
Exercise 7.3: Write a String method for the *tree type in gopl.io/ch4/treesort
(§4.4) that reveals the sequence of values in the tree.
*/

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import (
	"fmt"
)

// !+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

var Add = add

func (t *tree) String() string {
	return fmt.Sprintf("[%s]", traverse(t))
}

func traverse(t *tree) string {
	if t != nil {
		var l, r string
		if ll := traverse(t.left); len(ll) > 0 {
			l = fmt.Sprintf("(%s)", ll)
		}

		if rr := traverse(t.right); len(rr) > 0 {
			r = fmt.Sprintf("(%s)", rr)
		}

		return fmt.Sprintf("%d%s%s", t.value, l, r)
	}
	return ""
}

//!-
