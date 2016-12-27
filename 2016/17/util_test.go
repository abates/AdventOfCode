package main

import (
	"reflect"
	"testing"
)

func TestNeighbors(t *testing.T) {
	tests := []struct {
		x      int
		y      int
		path   string
		result []string
	}{
		{1, 1, "hijkl", []string{"hijklD"}},
		{1, 2, "hijklD", []string{"hijklDU", "hijklDR"}},
	}

	for i, test := range tests {
		v := NewVaultCoordinate(test.x, test.y, test.path)
		result := []string{}
		for _, node := range v.Neighbors() {
			result = append(result, node.ID())
		}
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("Test %d expected %v got %v", i, test.result, result)
		}
	}
}
