package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"0 3  0  1  -3", []int{0, 3, 0, 1, -3}},
	}

	for i, test := range tests {
		values := parse(test.input)
		if !reflect.DeepEqual(test.expected, values) {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, values)
		}
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0 3 0 1 -3", "2 3 2 3 -1"},
		{"0 1 2 3 4", "2 2 3 3 3"},
	}

	for i, test := range tests {
		values := parse(test.input)
		expected := parse(test.expected)
		computer := NewComputer(values)
		computer.Execute()
		if !reflect.DeepEqual(expected, values) {
			t.Errorf("tests[%d] expected %-v got %-v", i, expected, values)
		}
	}
}
