package alg

import (
	"testing"
)

func TestTraverse(t *testing.T) {
	tests := []struct {
		startX      int
		startY      int
		destination *Coordinate
		result      int
	}{
		{1, 1, NewCoordinate(4, 4), 7},
	}

	for i, test := range tests {
		var result []Node
		Traverse(NewCoordinate(test.startX, test.startY), func(l int, path []Node) bool {
			result = path
			return result[len(result)-1].(*Coordinate).Equal(test.destination)
		})

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
		result := TraverseLevel(NewCoordinate(1, 1), test.level)
		if len(result) != test.result {
			t.Errorf("Test %d: Unexpected path length.  Expected %d got %d %v", i, test.result, len(result), result)
		}
	}
}
