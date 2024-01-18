package main

import (
	"testing"
)

func TestStrJoin(t *testing.T) {

	if strJoin(", ", "1", "2", "3", "4") != "1, 2, 3, 4" {
		t.Error(`strJoin(",", "1", "2", "3", "4") != "1, 2, 3, 4"`)
	}

	if strJoin(", ") != "" {
		t.Error(`strJoin(", ") != ""`)
	}

	if strJoin(", ", "1") != "1" {
		t.Error(`strJoin(", ", "1") != "1"`)
	}

	if strJoin(" ", "1", "2", "3", "4") != "1 2 3 4" {
		t.Error(`strJoin(" ", "1", "2", "3", "4") != "1 2 3 4"`)
	}

	if strJoin("", "1", "2", "3", "4") != "1234" {
		t.Error(`strJoin(" ", "1", "2", "3", "4") != "1234"`)
	}
}
