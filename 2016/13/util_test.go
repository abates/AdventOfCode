package main

import (
	"github.com/abates/AdventOfCode/2016/alg"
	"testing"
)

func TestMagicDetector(t *testing.T) {
	tests := []struct {
		x      int
		y      int
		result bool
	}{
		{0, 0, true},
		{1, 0, false},
		{2, 0, true},
		{0, 1, true},
		{1, 1, true},
		{2, 1, false},
	}

	detector := MagicDetector(10)
	for i, test := range tests {
		result := detector(&alg.Coordinate{test.x, test.y})
		if result != test.result {
			t.Errorf("Test %d failed.  Expected %v got %v", i, result, test.result)
		}
	}
}

func TestWalk(t *testing.T) {
	tests := []struct {
		origin      *alg.Coordinate
		destination *alg.Coordinate
		result      int
		walkFunc    WallDetector
	}{
		{&alg.Coordinate{1, 1}, &alg.Coordinate{7, 4}, 12, MagicDetector(10)},
		{&alg.Coordinate{1, 1}, &alg.Coordinate{4, 4}, 7, func(*alg.Coordinate) bool { return true }},
	}

	for i, test := range tests {
		result := Walk(test.origin, test.destination, test.walkFunc)
		if result != test.result {
			t.Errorf("Test %d failed.  Expected %d Got %d", i, test.result, result)
		}
	}
}

func TestWalkMax(t *testing.T) {
	tests := []struct {
		max    int
		result int
	}{
		{1, 5},
	}
	for i, test := range tests {
		result := WalkMax(&alg.Coordinate{1, 1}, func(*alg.Coordinate) bool { return true }, test.max)
		if result != test.result {
			t.Errorf("Test %d failed.  Expected %d got %d", i, test.result, result)
		}
	}
}

func TestBranch(t *testing.T) {
	expectedDirections := []*MazeCoordinate{&MazeCoordinate{&alg.Coordinate{1, 2}, nil}, &MazeCoordinate{&alg.Coordinate{0, 1}, nil}}
	coordinate := &MazeCoordinate{&alg.Coordinate{1, 1}, MagicDetector(10)}
	directions := coordinate.Neighbors()

	for i, direction := range directions {
		if !expectedDirections[i].Equal(direction) {
			t.Errorf("Test %d Expected coordinate %s got %s", i, expectedDirections[i], direction)
		}
	}
}

func TestDraw(t *testing.T) {
	expected := "  0123456789\n0 .#.####.##\n1 ..#..#...#\n2 #....##...\n3 ###.#.###.\n4 .##..#..#.\n5 ..##....#.\n6 #...##.###\n"
	result := Draw(10, 7, MagicDetector(10))
	if expected != result {
		t.Errorf("Expected\n%sGot\n%s", expected, result)
	}
}
