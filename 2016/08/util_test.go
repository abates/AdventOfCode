package main

import (
	"testing"
)

type TestValues struct {
	width  int
	height int
	result string
}

func TestString(t *testing.T) {
	tests := []struct {
		width  int
		height int
		result string
	}{
		{5, 1, "     "},
		{5, 2, "     \n     "},
		{5, 3, "     \n     \n     "},
	}

	for i, test := range tests {
		s := NewScreen(test.width, test.height)
		result := s.String()
		if result != test.result {
			t.Errorf("Test %d failed.  Expected:\n%s\nReceived:\n%s\n", i, test.result, result)
		}
	}
}

func TestRect(t *testing.T) {
	tests := []struct {
		width      int
		height     int
		rectWidth  int
		rectHeight int
		result     string
	}{
		{5, 2, 2, 2, "##   \n##   "},
		{5, 5, 3, 3, "###  \n###  \n###  \n     \n     "},
		{5, 5, 6, 6, "#####\n#####\n#####\n#####\n#####"},
	}

	for i, test := range tests {
		s := NewScreen(test.width, test.height)
		s.Rect(test.rectWidth, test.rectHeight)
		result := s.String()
		if result != test.result {
			t.Errorf("Test %d failed.  Expected:\n%s\nReceived:\n%s\n", i, test.result, result)
		}
	}
}

func TestRotateRow(t *testing.T) {
	tests := []struct {
		width        int
		height       int
		rectWidth    int
		rectHeight   int
		rotateRow    int
		rotateAmount int
		result       string
	}{
		{5, 2, 2, 2, 1, 1, "##   \n ##  "},
		{5, 2, 2, 2, 1, 4, "##   \n#   #"},
	}

	for i, test := range tests {
		s := NewScreen(test.width, test.height)
		s.Rect(test.rectWidth, test.rectHeight)
		s.RotateRow(test.rotateRow, test.rotateAmount)
		result := s.String()
		if result != test.result {
			t.Errorf("Test %d failed.  Expected:\n%s\nReceived:\n%s\n", i, test.result, result)
		}
	}
}

func TestRotateColumn(t *testing.T) {
	tests := []struct {
		width        int
		height       int
		rectWidth    int
		rectHeight   int
		rotateColumn int
		rotateAmount int
		result       string
	}{
		{5, 3, 2, 2, 1, 1, "#    \n##   \n #   "},
	}

	for i, test := range tests {
		s := NewScreen(test.width, test.height)
		s.Rect(test.rectWidth, test.rectHeight)
		s.RotateColumn(test.rotateColumn, test.rotateAmount)
		result := s.String()
		if result != test.result {
			t.Errorf("Test %d failed.  Expected:\n%s\nReceived:\n%s\n", i, test.result, result)
		}
	}
}
