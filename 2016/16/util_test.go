package main

import (
	"math/big"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		bitLen int
		a      int64
		result int64
	}{
		{1, 0x01, 0x04},
		{5, 0x1f, 0x7c0},
		{12, 0xf0a, 0x1e14af0},
		{2, 0x01, 0x09},
		{9, 0xd5, 0x354a9},
	}

	for i, test := range tests {
		expected := big.NewInt(test.result)
		a := big.NewInt(test.a)
		a = Generate(test.bitLen, a)

		if a.Cmp(expected) != 0 {
			t.Errorf("Test %d Expected %b got %b", i, expected, a)
		}
	}
}

func TestFill(t *testing.T) {
	i := big.NewInt(0x10)
	a := Fill(20, i.BitLen(), i)
	expected := big.NewInt(0x83c87)
	if a.Cmp(expected) != 0 {
		t.Errorf("Expected %b got %b", expected, a)
	}
}

func TestChecksum(t *testing.T) {
	tests := []struct {
		fillLen int
		input   string
		result  string
	}{
		{20, "10000", "01100"},
		{10, "011010101", "00000"},
	}

	for i, test := range tests {
		a, _ := FillString(test.fillLen, test.input)
		result := Checksum(test.input[0] == '0', a)
		if result != test.result {
			t.Errorf("Test %d failed. Expected %s got %s", i, test.result, result)
		}
		/*i := big.NewInt(0x10)
		a := Fill(20, i.BitLen(), i)
		checksum := Checksum(false, a)
		expected := "01100"

		if checksum != expected {
			t.Errorf("Expected %s got %s", expected, checksum)
		}*/
	}
}
