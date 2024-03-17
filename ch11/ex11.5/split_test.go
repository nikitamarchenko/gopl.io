/*

Exercise 11.5: Extend TestSplit to use a table of inputs and expected outputs.

*/

package split

import (
	"reflect"
	"strings"
	"testing"
)

func TestSpit(t *testing.T) {
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"simple test", args{"a:b:c", ":"}, []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Spit() = %v, want %v", got, tt.want)
			}
		})
	}
}
