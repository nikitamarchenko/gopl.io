/*
Exercise 7.17: Extend xmlselect so that elements may be selected not just by
name, but by their attributes too, in the manner of CSS, so that, for instance,
an element like <div id="page" class="wide"> could be selected by a matching
id or class as well as its name.
*/

package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	type args struct {
		t Token
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantR   bool
		wantErr bool
	}{
		{"div", args{t: Token{name: "div"}, s: "div"}, true, false},
		{"div1", args{t: Token{name: "div"}, s: "div1"}, false, false},
		{"div[]", args{t: Token{name: "div"}, s: "div[]"}, true, false},
		{"div[class]", args{Token{"div", Attrs{"class":""}}, "div[class]" }, true, false},
		{"div[c1ass]", args{Token{"div", Attrs{"class":""}}, "div[c1ass]" }, false, false},
		{"div[class=name]", args{Token{"div", Attrs{"class":"name"}}, "div[class=name]" }, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, err := Check(tt.args.t, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotR != tt.wantR {
				t.Errorf("Check() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
