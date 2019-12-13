package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestD6Parse(t *testing.T) {
	COM := &orbitNode{name: "COM"}
	B := &orbitNode{name: "B", center: COM}
	C := &orbitNode{name: "C", center: B}
	COM.orbits = append(COM.orbits, B)
	B.orbits = append(B.orbits, C)

	tests := []struct {
		name  string
		input string
		want  *orbitNode
	}{
		{"test 1", "COM)B\nB)C", COM},
	}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			d6 := &D6{}
			for _, line := range strings.Split(test.input, "\n") {
				err := d6.parse(line)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			got := d6.orbitMap
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("Wanted %v got %v\n", test.want, got)
			}
		})
	}
}

func TestD6Parts(t *testing.T) {
	tests := []challengeTest{
		{"test 1", "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L", "IO Count: 42", ""},
		{"test 2", "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN", "", "4 orbital transfers required"},
	}

	for _, test := range tests {
		d6 := &D6{}
		challenge := &challenge{"Test Day 06", "", d6.parse, nil, d6.part1, d6.part2}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}
