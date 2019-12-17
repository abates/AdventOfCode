package main

import "testing"

func TestD13Parse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			d13 := &D13{}
			err := parseFile(test.input, &challenge{"Test Day 13", "", d13})
			if err == nil {
			} else {
				t.Errorf("Unexpected Error: %v", err)
			}
		})
	}
}

func TestD13Parts(t *testing.T) {
	tests := []challengeTest{}

	for _, test := range tests {
		d13 := &D13{}
		challenge := &challenge{"Test Day 13", "", d13}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}
