package main

import "testing"

func TestBit(t *testing.T) {
	tests := []struct {
		input    []byte
		index    int
		expected int
	}{
		{[]byte{0x00, 0x00, 0x01}, 23, 1},
		{[]byte{0x00, 0x00, 0x01}, 0, 0},
		{[]byte{0x80, 0x00, 0x00}, 0, 1},
		{[]byte{0x80, 0x00, 0x00}, 23, 0},
		{[]byte{0x80, 0x40, 0x00}, 9, 1},
	}

	for i, test := range tests {
		row := &Row{test.input}
		v := row.Bit(test.index)
		if test.expected != v {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, v)
		}
	}
}

func TestDiskCount(t *testing.T) {
	tests := []struct {
		input    []byte
		expected int
	}{
		{[]byte{0x00, 0x01, 0xd4}, 5},
	}

	for i, test := range tests {
		d := NewDisk()
		d.Append(test.input)
		c := d.Count()
		if test.expected != c {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, c)
		}
	}
}

func TestDiskGroups(t *testing.T) {
	tests := []struct {
		input    [][]byte
		expected int
	}{
		{
			input: [][]byte{
				[]byte{0xd4, 0x11},
				[]byte{0x55, 0x11},
				[]byte{0x0a, 0x11},
				[]byte{0xad, 0x11},
			},
			expected: 11,
		},
	}

	for i, test := range tests {
		d := NewDisk()
		for _, buf := range test.input {
			d.Append(buf)
		}
		c := d.Groups()
		if test.expected != c {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, c)
		}
	}
}
