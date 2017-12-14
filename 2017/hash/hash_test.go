package hash

import (
	"reflect"
	"testing"
)

func TestGetSet(t *testing.T) {
	input := Hash{0, 1, 2, 3, 4, 5}
	tests := []struct {
		index    int
		expected byte
	}{
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
		{5, 5},
		{6, 5},
		{7, 6},
	}

	for i, test := range tests {
		value := input.get(i)
		if test.expected != value {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, value)
		}

		input.set(i, value+5)
	}

	expected := []byte{10, 11, 7, 8, 9, 10}
	if !reflect.DeepEqual(expected, []byte(input)) {
		t.Errorf("expected %+v got %+v", expected, []byte(input))
	}
}

func TestReverse(t *testing.T) {
	list := Hash{0, 1, 2, 3, 4}
	tests := []struct {
		start    int
		end      int
		expected []byte
	}{
		// lengths 3, 4, 1, 5
		{0, 3, []byte{2, 1, 0, 3, 4}},   // skip size 0
		{3, 7, []byte{4, 3, 0, 1, 2}},   // skip size 1
		{8, 1, []byte{4, 3, 0, 1, 2}},   // skip size 2
		{11, 15, []byte{4, 2, 1, 0, 3}}, // skip size 3
	}

	for i, test := range tests {
		list.reverse(test.start, test.end)
		if !reflect.DeepEqual(test.expected, []byte(list)) {
			t.Errorf("tests[%d] expected %+v got %v", i, test.expected, []byte(list))
		}
	}
}

func TestXor(t *testing.T) {
	tests := []struct {
		input    []byte
		expected byte
	}{
		{[]byte{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}, 64},
	}

	for i, test := range tests {
		value := xor(test.input)
		if test.expected != value {
			t.Errorf("tests[%d] expected 0x%02x got 0x%02x", i, test.expected, value)
		}
	}
}

func TestComputeString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	for i, test := range tests {
		hash := ComputeString(test.input)
		str := hash.String()
		if str != test.expected {
			t.Errorf("tests[%d] expected %q got %q", i, test.expected, str)
		}
	}
}
