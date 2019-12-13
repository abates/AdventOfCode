package main

import (
	"reflect"
	"testing"
)

func TestD8Parse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{"test 1", "12345", []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			d8 := &D8{}
			err := d8.parse(test.input)
			if err == nil {
				if !reflect.DeepEqual(test.want, d8.input) {
					t.Errorf("Wanted %v got %v", test.want, d8.input)
				}
			} else {
				t.Errorf("Unexpected Error: %v", err)
			}
		})
	}
}

func TestD8ImageParse(t *testing.T) {
	tl := func(width, height int, rows [][]int) *Layer {
		return &Layer{width: width, height: height, rows: rows}
	}

	tests := []struct {
		name   string
		input  []int
		width  int
		height int
		want   []*Layer
	}{
		{"test 1", []int{1, 2, 3, 4, 5, 6}, 3, 2, []*Layer{tl(3, 2, [][]int{{1, 2, 3}, {4, 5, 6}})}},
		{"test 1", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, 3, 2, []*Layer{tl(3, 2, [][]int{{1, 2, 3}, {4, 5, 6}}), tl(3, 2, [][]int{{7, 8, 9}, {0, 1, 2}})}},
	}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			img := &Image{width: test.width, height: test.height}
			err := img.Parse(test.input)
			if err == nil {
				for _, layer := range img.layers {
					layer.index = nil
				}
				if !reflect.DeepEqual(test.want, img.layers) {
					t.Errorf("Wanted %v got %v", test.want, img.layers)
				}
			} else {
				t.Errorf("Unexpected Error: %v", err)
			}
		})
	}
}

func TestD8Parts(t *testing.T) {
	tests := []challengeTest{
		{"test 1", "0222112222120000", "", "\n *\n* \n"},
	}

	for _, test := range tests {
		d8 := &D8{width: 2, height: 2}
		challenge := &challenge{"Test Day 08", "", d8.parse, nil, d8.part1, d8.part2}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}
