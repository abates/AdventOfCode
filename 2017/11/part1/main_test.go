package main

import "testing"

func TestDistance(t *testing.T) {
	tests := []struct {
		input    string
		distance int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	}

	for i, test := range tests {
		m, _ := positionChild(test.input)
		distance := distance(m)
		if distance != test.distance {
			t.Errorf("tests[%d] expected %d got %d", i, test.distance, distance)
		}
	}
}
