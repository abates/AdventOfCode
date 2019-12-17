package main

import (
	"testing"
)

func TestD1Parse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{
		{"Mass 12", "12", 12},
		{"Mass 14", "14", 14},
		{"Mass 1969", "1969", 1969},
		{"Mass 100756", "100756", 100756},
	}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			d1 := &D1{}
			err := d1.parse(test.input)
			if err == nil {
				if len(d1.modules) != 1 {
					t.Errorf("Expected 1 module to be added, got %d", len(d1.modules))
				} else if d1.modules[0] != test.want {
					t.Errorf("Wanted module mass %d got %d", test.want, d1.modules[0])
				}
			} else {
				t.Errorf("Unexpected Error: %v", err)
			}
		})
	}
}

func TestD1Parts(t *testing.T) {
	tests := []challengeTest{
		{"", "12\n14\n1969\n100756", "Sum of fuel requirements is 34241", "Sum of fuel requirements is 51316"},
	}

	for _, test := range tests {
		d1 := &D1{}
		challenge := &challenge{"Test Day 01", "", d1}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}

func TestMFR(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Mass 12", 12, 2},
		{"Mass 14", 14, 2},
		{"Mass 1969", 1969, 654},
		{"Mass 100756", 100756, 33583},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := MFR(test.input)
			if test.want != got {
				t.Errorf("Wanted fuel requirement %d got %d", test.want, got)
			}
		})
	}
}

func TestTotalMFR(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"Mass 14", 14, 2},
		{"Mass 1969", 1969, 966},
		{"Mass 100756", 100756, 50346},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TotalMFR(test.input)
			if test.want != got {
				t.Errorf("Wanted fuel requirement %d got %d", test.want, got)
			}
		})
	}
}
