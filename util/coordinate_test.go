package util

import (
	"fmt"
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		x1       int
		y1       int
		x2       int
		y2       int
		expected int
	}{
		{0, 0, 1, 1, 2},
		{0, 0, 2, 2, 4},
		{-1, -1, 1, 1, 4},
		{1, 1, 1, 1, 0},
	}

	for i, test := range tests {
		start := NewCoordinate(test.x1, test.y1)
		end := NewCoordinate(test.x2, test.y2)

		d := ManhattanDistance(start, end)
		if d != test.expected {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, d)
		}

		d = ManhattanDistance(start, end)
		if d != test.expected {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, d)
		}
	}
}

func TestAddSubtractCoordinates(t *testing.T) {
	tests := []struct {
		x     int
		y     int
		diffX int
		diffY int
		addX  int
		addY  int
		subX  int
		subY  int
	}{
		{1, 1, 0, -1, 1, 0, 1, 2},
		{1, 1, 0, -1, 1, 0, 1, 2},
		{1, 1, 0, -1, 1, 0, 1, 2},
		{1, 1, -1, 0, 0, 1, 2, 1},
	}

	for i, test := range tests {
		coord := NewCoordinate(test.x, test.y)
		diffX := NewCoordinate(test.diffX, test.diffY)

		result := coord.Add(diffX)
		expected := NewCoordinate(test.addX, test.addY)

		str := fmt.Sprintf("(%d,%d)", test.x, test.y)
		if str != coord.String() {
			t.Errorf("tests[%d] expected %s got %s", i, str, coord.String())
		}

		if !result.Equal(expected) {
			t.Errorf("tests[%d] Add expected %s got %s", i, expected, result)
		}

		result = coord.Subtract(diffX)
		expected = NewCoordinate(test.subX, test.subY)

		if !result.Equal(expected) {
			t.Errorf("tests[%d] Subtract expected %s got %s", i, expected, result)
		}
	}
}

/*
func TestNeighborsDiagonals(t *testing.T) {
	coordinate := NewCoordinate(0, 0)
	expectedNeighbors := [4]*Coordinate{
		NewCoordinate(0, -1),
		NewCoordinate(1, 0),
		NewCoordinate(0, 1),
		NewCoordinate(-1, 0),
	}

	var neighbors [4]Coordinate
	i := 0
	coordinate.Neighbors(func(c *Coordinate) {
		neighbors[i] = c
		i++
	})

	if !reflect.DeepEqual(expectedNeighbors, neighbors) {
		t.Errorf("Test expected %-v got %-v", expectedNeighbors, neighbors)
	}

	expectedDiagonals := [4]*Coordinate{
		NewCoordinate(-1, 1),
		NewCoordinate(1, 1),
		NewCoordinate(1, -1),
		NewCoordinate(-1, -1),
	}

	var diagonals [4]Coordinate
	i = 0
	coordinate.Diagonals(func(c *Coordinate) {
		diagonals[i] = c
		i++
	})

	if !reflect.DeepEqual(expectedDiagonals, diagonals) {
		t.Errorf("Test expected %-v got %-v", expectedDiagonals, diagonals)
	}
}*/
