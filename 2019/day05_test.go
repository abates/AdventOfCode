package main

import "testing"

func TestD5Parts(t *testing.T) {
	tests := []challengeTest{}

	for _, test := range tests {
		d5 := &D5{}
		challenge := &challenge{"Test Day 05", "", nil, d5.parseFile, d5.part1, d5.part2}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}
