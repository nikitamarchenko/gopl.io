/*

Exercise 13.2: Write a function that reports whether its argument is a cyclic data structure.

*/

package cycle

import (
	"testing"
)

func TestIsCycle(t *testing.T) {

	type sliceT []*sliceT
	var slice sliceT
	slice = append(slice, &slice)

	type structT struct {
		s *structT 
	}
    var strct structT
	strct.s = &strct

	type mapT map[int]*mapT
	mm := make(mapT)
	mm[0] = &mm

	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"int", args{1}, false},
		{"string", args{"42"}, false},
		{"slice", args{sliceT{}}, false},
		{"slice cycle", args{slice}, true},
		{"struct", args{structT{}}, false},
		{"struct cycle", args{strct}, true},
		{"map", args{make(mapT)}, false},
		{"map cycle", args{mm}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCycle(tt.args.x); got != tt.want {
				t.Errorf("IsCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
