package main

import (
	"github.com/abates/AdventOfCode/2016/util"
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
		result := detector(&util.Coordinate{test.x, test.y})
		if result != test.result {
			t.Errorf("Test %d failed.  Expected %v got %v", i, result, test.result)
		}
	}
}

func TestWalk(t *testing.T) {
	tests := []struct {
		startX      int
		startY      int
		destination *util.Coordinate
		result      int
		walkFunc    WallDetector
	}{
		{1, 1, &util.Coordinate{7, 4}, 12, MagicDetector(10)},
		{1, 1, &util.Coordinate{4, 4}, 7, func(*util.Coordinate) bool { return true }},
	}

	for i, test := range tests {
		walker := NewWalker(test.startX, test.startY, test.walkFunc)
		result := walker.Walk(test.destination)
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
		walker := NewWalker(1, 1, func(*util.Coordinate) bool { return true })
		result := walker.WalkMax(test.max)
		if result != test.result {
			t.Errorf("Test %d failed.  Expected %d got %d", i, test.result, result)
		}
	}
}

func TestBranch(t *testing.T) {
	expectedDirections := []*MazeCoordinate{&MazeCoordinate{&util.Coordinate{1, 2}, nil}, &MazeCoordinate{&util.Coordinate{0, 1}, nil}}
	coordinate := &MazeCoordinate{&util.Coordinate{1, 1}, MagicDetector(10)}
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
