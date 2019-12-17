package main

import "testing"

func TestD5Parts(t *testing.T) {
	tests := []challengeTest{}

	for _, test := range tests {
		d5 := &D5{}
		challenge := &challenge{"Test Day 05", "", d5}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}
