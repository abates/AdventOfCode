package main

import (
	"reflect"
	"testing"
)

func TestGetSet(t *testing.T) {
	input := List{0, 1, 2, 3, 4, 5}
	tests := []struct {
		index    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 5},
		{7, 6},
	}

	for i, test := range tests {
		value := input.Get(i)
		if test.expected != value {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, value)
		}

		input.Set(i, value+5)
	}

	expected := []int{10, 11, 7, 8, 9, 10}
	if !reflect.DeepEqual(expected, []int(input)) {
		t.Errorf("expected %+v got %+v", expected, []int(input))
	}
}

func TestReverse(t *testing.T) {
	list := List{0, 1, 2, 3, 4}
	tests := []struct {
		start    int
		end      int
		expected []int
	}{
		// lengths 3, 4, 1, 5
		{0, 3, []int{2, 1, 0, 3, 4}},   // skip size 0
		{3, 7, []int{4, 3, 0, 1, 2}},   // skip size 1
		{8, 1, []int{4, 3, 0, 1, 2}},   // skip size 2
		{11, 15, []int{4, 2, 1, 0, 3}}, // skip size 3
	}

	for i, test := range tests {
		list.Reverse(test.start, test.end)
		if !reflect.DeepEqual(test.expected, []int(list)) {
			t.Errorf("tests[%d] expected %+v got %v", i, test.expected, []int(list))
		}
	}
}

func TestCompute(t *testing.T) {
	input := List{0, 1, 2, 3, 4}
	lengths := []int{3, 4, 1, 5}
	expected := List{3, 4, 2, 1, 0}
	compute(input, lengths)

	if !reflect.DeepEqual(expected, input) {
		t.Errorf("expected %+v got %+v", expected, input)
	}
}
