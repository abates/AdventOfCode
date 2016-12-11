package main

import (
	"strings"
	"testing"
)

func TestV1Decompress(t *testing.T) {
	tests := []struct {
		input  string
		result string
	}{
		{"ADVENT", "ADVENT"},
		{"A(1x5)BC", "ABBBBBC"},
		{"(3x3)XYZ", "XYZXYZXYZ"},
		{"A(2x2)BCD(2x2)EFG", "ABCBCDEFEFG"},
		{"(6x1)(1x3)A", "(1x3)A"},
		{"(6x2)(1x3)A", "(1x3)A(1x3)A"},
		{"X(8x2)(3x3)ABCY", "X(3x3)ABC(3x3)ABCY"},
		{"X(1x5)BC(1x5)D", "XBBBBBCDDDDD"},
		{"XXXXXXXXXX(1x5)BC(1x5)D", "XXXXXXXXXXBBBBBCDDDDD"},
	}

	for i, test := range tests {
		result := V1EncodedString(test.input).Decode()
		if result != len(test.result) {
			t.Errorf("Test %d failed.  Expected %d got %d", i, len(test.result), result)
		}
	}
}

func TestV2Decompress(t *testing.T) {
	tests := []struct {
		input  string
		result string
	}{
		{"(3x3)XYZ", "XYZXYZXYZ"},
		{"X(8x2)(3x3)ABCY", "XABCABCABCABCABCABCY"},
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", strings.Repeat("A", 241920)},
	}

	for i, test := range tests {
		result := V2EncodedString(test.input).Decode()
		if result != len(test.result) {
			t.Errorf("Test %d failed.  Expected %d got %d", i, len(test.result), result)
		}
	}
}
