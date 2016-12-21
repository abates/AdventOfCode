package main

import (
	"testing"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		input  string
		x      int
		y      int
		result string
	}{
		{"abcde", 4, 0, "ebcda"},
	}

	for i, test := range tests {
		password := NewPassword(test.input)
		password.Swap(test.x, test.y)
		if password.String() != test.result {
			t.Errorf("Test %d expected %s got %s", i, test.result, password.String())
		}
	}
}

func TestSwapLetter(t *testing.T) {
	tests := []struct {
		input  string
		x      string
		y      string
		result string
	}{
		{"ebcda", "b", "d", "edcba"},
	}
	for i, test := range tests {
		password := NewPassword(test.input)
		password.SwapLetter(test.x, test.y)
		if password.String() != test.result {
			t.Errorf("Test %d expected %s got %s", i, test.result, password.String())
		}
	}

}

func TestRotate(t *testing.T) {
	tests := []struct {
		input     string
		x         int
		direction string
		result    string
	}{
		{"ebcda", 1, "LEFT", "bcdae"},
		{"ebcda", 6, "LEFT", "bcdae"},
		{"bcdea", 1, "RIGHT", "abcde"},
		{"bcdea", 6, "RIGHT", "abcde"},
	}

	for i, test := range tests {
		password := NewPassword(test.input)
		if test.direction == "LEFT" {
			password.RotateLeft(test.x)
		} else if test.direction == "RIGHT" {
			password.RotateRight(test.x)
		} else {
			panic("Unknown direction " + test.direction)
		}

		if password.String() != test.result {
			t.Errorf("Test %d expected %s got %s", i, test.result, password.String())
		}
	}
}

func TestRotatePosition(t *testing.T) {
	tests := []struct {
		input     string
		direction string
		x         string
		result    string
	}{
		{"abdec", "b", "ecabd"},
		{"ecabd", "d", "decab"},
	}

	for i, test := range tests {
		password := NewPassword(test.input)
		password.RotatePosition(test.x)
		if password.String() != test.result {
			t.Errorf("Test %d expected %s got %s", i, test.result, password.String())
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input  string
		x      int
		y      int
		result string
	}{
		{"edcba", 0, 4, "abcde"},
	}

	for i, test := range tests {
		password := NewPassword(test.input)
		password.Reverse(test.x, test.y)
		if password.String() != test.result {
			t.Errorf("Test %d expected %s got %s", i, test.result, password.String())
		}
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		input  string
		x      int
		y      int
		result string
	}{
		{"bcdea", 1, 4, "bdeac"},
		{"bdeac", 3, 0, "abdec"},
	}

	for i, test := range tests {
		password := NewPassword(test.input)
		password.Move(test.x, test.y)
		if password.String() != test.result {
			t.Errorf("Test %d expected %s got %s", i, test.result, password.String())
		}
	}
}
