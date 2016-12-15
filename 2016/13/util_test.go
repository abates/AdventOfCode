package main

import (
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
		result := detector(&Coordinate{test.x, test.y})
		if result != test.result {
			t.Errorf("Test %d failed.  Expected %v got %v", i, result, test.result)
		}
	}
}

func TestWalk(t *testing.T) {
	walker := NewMagicWalker(1, 1, 10)
	result := walker.Walk(Coordinate{7, 4})
	if result != 11 {
		t.Errorf("Test failed.  Expected %d Got %d", 11, result)
	}
}

func TestWalk1(t *testing.T) {
	walker := NewWalker(1, 1, func(*Coordinate) bool { return true })
	result := walker.Walk(Coordinate{4, 4})
	if result != 6 {
		t.Errorf("Test failed.  Expected %d got %d", 6, result)
	}
}

func TestWalkMax(t *testing.T) {
	tests := []struct {
		destination Coordinate
		max         int
		result      int
	}{
		{Coordinate{4, 4}, 1, 5},
	}
	for i, test := range tests {
		walker := NewWalker(1, 1, func(*Coordinate) bool { return true })
		result := walker.WalkMax(test.destination, test.max)
		if result != test.result {
			t.Errorf("Test %d failed.  Expected %d got %d", i, test.result, result)
		}
	}
}

func TestAdd(t *testing.T) {
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

func TestBranch(t *testing.T) {
	expectedDirections := []Coordinate{{1, 2}, {0, 1}}
	walker := NewMagicWalker(1, 1, 10)
	directions := walker.Next()

	for i, direction := range directions {
		if !expectedDirections[i].Equal(direction) {
			t.Errorf("Expected coordinate %s got %s", expectedDirections[i], direction)
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
