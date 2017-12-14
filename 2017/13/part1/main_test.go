package main

import "testing"

func TestPosition(t *testing.T) {
	tests := []struct {
		layer    Layer
		expected int
	}{
		// depth, range, step
		{Layer{0, 3, 0}, 0},
		{Layer{0, 3, 1}, 1},
		{Layer{0, 3, 2}, 2},
		{Layer{0, 3, 3}, 1},
		{Layer{0, 3, 4}, 0},
		{Layer{0, 3, 5}, 1},
		{Layer{0, 3, 6}, 2},
		{Layer{0, 3, 7}, 1},
		{Layer{0, 3, 8}, 0},
		{Layer{0, 3, 9}, 1},
	}

	for i, test := range tests {
		position := test.layer.Position()
		if position != test.expected {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, position)
		}
	}
}

func TestPenalize(t *testing.T) {
	tests := []struct {
		input           string
		expectedPenalty int
		expectedDelay   int
	}{
		{"0: 3\n1: 2\n4: 4\n6: 4\n", 24, 10},
	}

	for i, test := range tests {
		fw := initializeFirewall(test.input)
		pkt := Packet(0)
		fw.Traverse(&pkt)
		if int(pkt) != test.expectedPenalty {
			t.Errorf("tests[%d] expected %d got %d", i, test.expectedPenalty, int(pkt))
		}

		delay := fw.TraversalDelay()
		if test.expectedDelay != delay {
			t.Errorf("tests[%d] expected %d got %d", i, test.expectedDelay, delay)
		}
	}
}
