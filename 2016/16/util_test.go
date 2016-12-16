package main

import (
	"bytes"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		bitLen int
		a      []byte
		result []byte
	}{
		{1, []byte{0x80}, []byte{0x80}},
		{5, []byte{0xf8, 00}, []byte{0xf8, 0x00}},
		{12, []byte{0xf0, 0xa0, 0x00, 0x00}, []byte{0xf0, 0xa5, 0x78, 0x00}},
		{2, []byte{0x80}, []byte{0x90}},
		{9, []byte{0x6a, 0x80, 0x00}, []byte{0x6a, 0x95, 0x20}},
	}

	for i, test := range tests {
		disk := &Disk{
			size: len(test.a) * 8,
			byts: test.a,
		}
		Generate(test.bitLen, disk)

		if !bytes.Equal(test.a, test.result) {
			t.Errorf("Test %d Expected %08b got %08b", i, test.result, test.a)
		}
	}
}

func TestFillString(t *testing.T) {
	disk := NewDisk(20)
	disk.Fill("10000")
	expected := []byte{0x83, 0xc8, 0x70}
	if !bytes.Equal(disk.byts, expected) {
		t.Errorf("Expected %08b got %08b", expected, disk.byts)
	}
}

func TestChecksum(t *testing.T) {
	tests := []struct {
		size   int
		input  string
		result string
	}{
		{20, "10000", "01100"},
		{10, "011010101", "00000"},
	}

	for i, test := range tests {
		disk := NewDisk(test.size)
		disk.Fill(test.input)

		result := disk.Checksum()
		if result != test.result {
			t.Errorf("Test %d failed. Expected %s got %s", i, test.result, result)
		}
	}
}
