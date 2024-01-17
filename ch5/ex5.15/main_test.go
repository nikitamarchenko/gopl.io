package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	if sum() != 0 {
		t.Error("sum() != 0")
	}

	if sum(3) != 3 {
		t.Error("sum() != 3")
	}

	if sum(1, 2, 3, 4) != 10 {
		t.Error("sum(1, 2, 3, 4) =! 10")
	}

	values := []int{1, 2, 3, 4}
	if sum(values...) != 10 {
		t.Error("values := []int{1, 2, 3, 4};sum(values...) != 10")
	}
}

func TestMin(t *testing.T) {

	if min(3) != 3 {
		t.Error("min() != 3")
	}

	if min(1, 2, 3, 4) != 1 {
		t.Error("min(1, 2, 3, 4) != 1")
	}

	values := []int{1, 2, 3, 4}
	if min(values...) != 1 {
		t.Error("values := []int{1, 2, 3, 4}; min(values...) != 1")
	}

	values = []int{4, 3, 2, 1}
	if min(values...) != 1 {
		t.Error("values := []int{4, 3, 2, 1}; min(values...) != 1")
	}
}

func TestMinNonEmpty(t *testing.T) {
	if minNonEmpty(1, 2, 3, 4) != 1 {
		t.Error("minNonEmpty(1, 2, 3, 4) != 1")
	}
}

func TestMax(t *testing.T) {

	if max(3) != 3 {
		t.Error("max() != 3")
	}

	if max(1, 2, 3, 4) != 4 {
		t.Error("max(1, 2, 3, 4) != 4")
	}

	values := []int{1, 2, 3, 4}
	if max(values...) != 4 {
		t.Error("values := []int{1, 2, 3, 4}; max(values...) != 4")
	}

	values = []int{4, 3, 2, 1}
	if max(values...) != 4 {
		t.Error("values := []int{4, 3, 2, 1}; max(values...) != 4")
	}
}

func TestMaxNonEmpty(t *testing.T) {
	if maxNonEmpty(1, 2, 3, 4) != 4 {
		t.Error("maxNonEmpty(1, 2, 3, 4) != 4")
	}
}
