package main

import (
	"reflect"
	"testing"
)

func TestRedistribute(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 2, 7, 0}, []int{2, 4, 1, 2}},
		{[]int{2, 4, 1, 2}, []int{3, 1, 2, 3}},
	}

	for i, test := range tests {
		mem := NewMemory(test.input)
		mem.Redistribute()
		values := make([]int, 0)
		for _, bank := range mem.banks {
			values = append(values, int(bank))
		}
		if !reflect.DeepEqual(test.expected, values) {
			t.Errorf("tests[%d] expected %-v got %-v", i, test.expected, values)
		}
	}
}

func TestCycle(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
		steps    int
	}{
		{[]int{0, 2, 7, 0}, []int{2, 4, 1, 2}, 5},
		{[]int{2, 4, 1, 2}, []int{2, 4, 1, 2}, 4},
	}

	for i, test := range tests {
		mem := NewMemory(test.input)
		steps := mem.Cycle()
		values := make([]int, 0)

		for _, bank := range mem.banks {
			values = append(values, int(bank))
		}

		if !reflect.DeepEqual(test.expected, values) {
			t.Errorf("tests[%d] expected %-v got %-v", i, test.expected, values)
		}

		if test.steps != steps {
			t.Errorf("tests[%d] expected %d got %d", i, test.steps, steps)
		}
	}
}
