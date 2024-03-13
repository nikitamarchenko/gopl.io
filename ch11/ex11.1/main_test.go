/*

Exercise 11.1: Write tests for the charcount program in Section 4.3.

*/

package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func Test_charcount(t *testing.T) {

	tests := []struct {
		name  string
		args  io.Reader
		want  map[rune]int
		want1 [utf8.UTFMax + 1]int
		want2 int
	}{
		{"test word charcount",
			strings.NewReader("charcount"),
			map[rune]int{'c': 2, 'h': 1, 'a': 1, 'r': 1, 'o': 1, 'u': 1, 'n': 1, 't': 1},
			[5]int{0, 9, 0, 0, 0},
			0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := charcount(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("charcount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("charcount() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("charcount() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
