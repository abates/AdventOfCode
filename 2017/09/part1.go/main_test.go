package main

import (
	"strings"
	"testing"
)

func TestParseGroup(t *testing.T) {
	tests := []struct {
		input string
		score int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}

	for i, test := range tests {
		group, _ := parseGroup(strings.NewReader(test.input))
		score := group.Score(0)
		if test.score != score {
			t.Errorf("tests[%d] expected %d got %d", i, test.score, score)
		}
	}
}

func TestCountGarbage(t *testing.T) {
	tests := []struct {
		input   string
		garbage int
	}{
		{"<>", 0},
		{"<random characters>", 17},
		{"<<<<>", 3},
		{"<{!>}>", 2},
		{"<!!>", 0},
		{"<!!!>>", 0},
		{"<{o\"i!a,<{i<a>", 10},
	}

	for i, test := range tests {
		_, garbage := parseGroup(strings.NewReader(test.input))
		if test.garbage != garbage {
			t.Errorf("tests[%d] expected %d got %d", i, test.garbage, garbage)
		}
	}
}
