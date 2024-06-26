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
	"math"
	"math/big"
	"reflect"
	"unsafe"
)

const threshold = 1e-9 // one part in a billion
func deepEqualFloat(x, y float64) bool {
	return math.Abs(x-y) < threshold*math.Max(x, y)
}

// we can't 
func deepEqualUint(x, y uint64) bool {
	var diff, max, z big.Float
	if x == y {
		return true
	} else if x > y {
		diff.SetUint64(x - y)
		max.SetUint64(x)
	} else {
		diff.SetUint64(y - x)
		max.SetUint64(y)
	}
	return diff.Cmp(z.Mul(big.NewFloat(threshold), &max)) == -1
}

func deepEqualInt(x, y int64) bool {
	return deepEqualFloat(float64(x), float64(y))
}

func deepEqualComplex(a, b complex128) bool {
	diffReal := real(a) - real(b)
	diffImag := imag(a) - imag(b)
	return diffReal < threshold && diffReal > -threshold && diffImag < threshold && diffImag > -threshold
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}
	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return deepEqualInt(x.Int(), y.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return deepEqualUint(x.Uint(), y.Uint())

	case reflect.Float32, reflect.Float64:
		return deepEqualFloat(x.Float(), y.Float())

	case reflect.Complex64, reflect.Complex128:
		return deepEqualComplex(x.Complex(), y.Complex())

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !equal(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}
