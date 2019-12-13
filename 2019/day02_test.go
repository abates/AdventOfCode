package main

import (
	"testing"
)

func TestD2ParseFile(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []int
	}{
		{"test 1", []string{"1,0,0,0,99", ""}, []int{1, 0, 0, 0, 99}},
		{"test 2", []string{"2,3,0,3,99", ""}, []int{2, 3, 0, 3, 99}},
		{"test 3", []string{"2,4,4,5,99,0", ""}, []int{2, 4, 4, 5, 99, 0}},
		{"test 4", []string{"1,1,1,4,99,5,6,0,99", ""}, []int{1, 1, 1, 4, 99, 5, 6, 0, 99}},
	}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			d2 := &D2{}
			err := d2.parseFile(test.input)
			if err == nil {
			} else {
				t.Errorf("Unexpected Error: %v", err)
			}
		})
	}
}
