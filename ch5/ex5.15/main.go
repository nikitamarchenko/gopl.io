/*
Exercise 5.15: Write variadic functions max and min, analogous to sum. What
should these functions do when called with no arguments? Write variants that
require at least one argument.
*/

package main

func sum(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func min(values ...int) int {
	if len(values) == 0 {
		panic("min: empty list")
	}

	r := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < r {
			r = values[i]
		}
	}
	return r
}

func minNonEmpty(value int, values ...int) int {

	tmp := make([]int, len(values)+1)
	tmp[0] = value
	cb := copy(tmp[1:], values)
	if cb != len(values) {
		panic("minNonEmpty: can't copy values")
	}
	return min(tmp...)
}

func max(values ...int) int {
	if len(values) == 0 {
		panic("min: empty list")
	}

	r := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > r {
			r = values[i]
		}
	}
	return r
}

func maxNonEmpty(value int, values ...int) int {

	tmp := make([]int, len(values)+1)
	tmp[0] = value
	cb := copy(tmp[1:], values)
	if cb != len(values) {
		panic("minNonEmpty: can't copy values")
	}
	return max(tmp...)
}
