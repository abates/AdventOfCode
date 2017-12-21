package main

import (
	"strings"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"12", "34"}, "2413"},
		{[]string{"123", "456", "789"}, "369258147"},
	}

	for i, test := range tests {
		b := NewBlock(test.input)
		b = b.Rotate()
		if b.String() != test.expected {
			t.Errorf("tests[%d] expected %q got %q", i, test.expected, b.String())
		}
	}
}

func TestTransform(t *testing.T) {
	tests := []struct {
		transforms []string
		input      []string
		expected   string
	}{
		{[]string{"../.# => ##./#../...", ".#./..#/### => #..#/..../..../#..#"}, []string{".#.", "..#", "###"}, "#..#........#..#"},
		{[]string{"../.# => ##./#../...", ".#./..#/### => #..#/..../..../#..#"}, []string{"#.", ".."}, "##.#....."},
		{[]string{"../.# => ##./#../...", ".#./..#/### => #..#/..../..../#..#"}, []string{".#", ".."}, "##.#....."},
		{[]string{"../.# => ##./#../...", ".#./..#/### => #..#/..../..../#..#"}, []string{"..", "#."}, "##.#....."},
		{[]string{"../.# => ##./#../...", ".#./..#/### => #..#/..../..../#..#"}, []string{"..", ".#"}, "##.#....."},
	}

	for i, test := range tests {
		transforms := NewTransforms()
		for _, t := range test.transforms {
			transforms.Add(t)
		}

		output := transforms.Transform(NewBlock(test.input))
		if output.String() != test.expected {
			t.Errorf("tests[%d] expected %q got %q", i, test.expected, output.String())
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		transforms []string
		input      []string
		expected   string
	}{
		{
			transforms: []string{"../.# => ##./#../...", ".#./..#/### => #..#/..../..../#..#"},
			input:      []string{".#.", "..#", "###"},
			expected:   "#..#........#..#",
		}, {
			transforms: []string{"../.# => ##./#../...", ".#./..#/### => #..#/..../..../#..#"},
			input:      []string{"#..#", "....", "....", "#..#"},
			expected:   "##.##.#..#........##.##.#..#........",
		},
	}

	for i, test := range tests {
		transforms := NewTransforms()
		for _, t := range test.transforms {
			transforms.Add(t)
		}
		grid := NewGrid(test.input)
		grid.Divide(transforms)

		output := strings.Join(grid.pixels, "")
		if output != test.expected {
			t.Errorf("tests[%d] expected %q got %q", i, test.expected, output)
		}
	}
}
