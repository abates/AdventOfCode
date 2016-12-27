package alg

import (
	"testing"
)

func TestAddCoordinates(t *testing.T) {
	tests := []struct {
		position *Coordinate
		addend   *Coordinate
		result   *Coordinate
	}{
		{&Coordinate{1, 1}, &Coordinate{0, -1}, &Coordinate{1, 0}},
		{&Coordinate{1, 1}, &Coordinate{1, 0}, &Coordinate{2, 1}},
		{&Coordinate{1, 1}, &Coordinate{0, -1}, &Coordinate{1, 0}},
		{&Coordinate{1, 1}, &Coordinate{-1, 0}, &Coordinate{0, 1}},
	}

	for i, test := range tests {
		result := test.position.Add(test.addend)
		if !result.Equal(test.result) {
			t.Errorf("Test %d failed.  Expected %s but got %s", i, test.result, result)
		}
	}
}
