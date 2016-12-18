package util

import (
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		startX      int
		startY      int
		destination *Coordinate
		result      int
	}{
		{1, 1, &Coordinate{4, 4}, 7},
	}

	for i, test := range tests {
		tree := NewTree(&Coordinate{1, 1})

		result := tree.Find(test.destination)
		if len(result) != test.result {
			t.Errorf("Test %d: Unexpected path length.  Expected %d got %d", i, test.result, len(result))
		}
	}
}

func TestFindAt(t *testing.T) {
	tests := []struct {
		startX int
		startY int
		level  int
		result int
	}{
		{1, 1, 1, 4},
		{1, 1, 2, 8},
		{1, 1, 3, 12},
		{1, 1, 4, 16},
	}

	for i, test := range tests {
		tree := NewTree(&Coordinate{1, 1})

		result := tree.FindAt(test.level)
		if len(result) != test.result {
			t.Errorf("Test %d: Unexpected path length.  Expected %d got %d %v", i, test.result, len(result), result)
		}
	}
}
