/*

Exercise 12.5: Adapt encode to emit JSON instead of S-expressions. Test your
encoder using the standard decoder, json.Unmarshal.

*/

package myjson

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshalNil(t *testing.T) {
	var arg *int
	got, err := Marshal(arg)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	var want *int
	err = json.Unmarshal(got, &want)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	if !reflect.DeepEqual(arg, want) {
		t.Errorf("Marshal() = %T %v, want %T %v", arg, arg, want, want)
	}
}

func TestMarshalInt(t *testing.T) {
	var arg int = 42
	got, err := Marshal(arg)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	var want int
	err = json.Unmarshal(got, &want)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	if !reflect.DeepEqual(arg, want) {
		t.Errorf("Marshal() = %T %v, want %T %v", arg, arg, want, want)
	}
}

func TestMarshalUInt(t *testing.T) {
	var arg uint = 42
	got, err := Marshal(arg)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	var want uint
	err = json.Unmarshal(got, &want)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	if !reflect.DeepEqual(arg, want) {
		t.Errorf("Marshal() = %T %v, want %T %v", arg, arg, want, want)
	}
}

func TestMarshalString(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
		{"empty", "", false},
		{"abcdefg", "abcdefg", false},
		{"escape", `\\ \/ \b \f \n \r \t \uAABBCCDD`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var want string
			err = json.Unmarshal(got, &want)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tt.arg, want) {
				t.Errorf("Marshal() = %v, want %v", tt.arg, want)
			}
		})
	}
}

func TestMarshalPtr(t *testing.T) {
	var uint42 uint = 42
	var arg *uint = &uint42

	got, err := Marshal(arg)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	var want *uint
	err = json.Unmarshal(got, &want)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
		return
	}
	if !reflect.DeepEqual(arg, want) {
		t.Errorf("Marshal() = %T %v, want %T %v", arg, arg, want, want)
	}
}

func TestMarshalSlice(t *testing.T) {
	arg := []string{"a", "bc", "def", "1 2 3", "456"}
	got, err := Marshal(arg)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	var want []string
	err = json.Unmarshal(got, &want)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
		return
	}
	if !reflect.DeepEqual(arg, want) {
		t.Errorf("Marshal() = %T %v, want %T %v", arg, arg, want, want)
	}
}

func TestMarshalObject(t *testing.T) {
	i42 := 42

	type S struct {
		A int
		B uint
		C string
		P *int
	}

	arg := S{1, 2, "test", &i42}

	got, err := Marshal(arg)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	var want S
	err = json.Unmarshal(got, &want)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
		return
	}
	if !reflect.DeepEqual(arg, want) {
		t.Errorf("Marshal() = %T %v, want %T %v", arg, arg, want, want)
	}
}

func TestMarshalMap(t *testing.T) {

	arg := make(map[string]int)
	arg["a"] = 1
	arg["b"] = 2
	arg["c"] = 3

	got, err := Marshal(arg)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	var want map[string]int
	err = json.Unmarshal(got, &want)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
		return
	}
	if !reflect.DeepEqual(arg, want) {
		t.Errorf("Marshal() = %T %v, want %T %v", arg, arg, want, want)
	}
}

func TestMarshalMapOfObjects(t *testing.T) {
	i42 := 42
	type S struct {
		A int
		B uint
		C string
		P *int
		M map[int]string
	}
	type M map[string]S

	arg := make(M)
	arg["a"] = S{1, 2, "test", &i42, map[int]string{1: "test2"}}

	got, err := Marshal(arg)
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
		return
	}
	var want M
	err = json.Unmarshal(got, &want)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
		return
	}
	if !reflect.DeepEqual(arg, want) {
		t.Errorf("Marshal() = %T %v, want %T %v", arg, arg, want, want)
	}
}
