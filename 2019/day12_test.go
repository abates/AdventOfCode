package main

import (
	"reflect"
	"testing"
)

func TestD12Parse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  OrbitalSystem
	}{
		{"test 1", "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>\n", OrbitalSystem{NewMoon(-1, 0, 2), NewMoon(2, -10, -7), NewMoon(4, -8, 8), NewMoon(3, 5, -1)}},
	}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			d12 := &D12{}
			err := parseFile(test.input, &challenge{"Test Day 12", "", d12})
			if err == nil {
				got := d12.moons
				if !reflect.DeepEqual(test.want, got) {
					t.Errorf("Wanted %v got %v", test.want, got)
				}
			} else {
				t.Errorf("Unexpected Error: %v", err)
			}
		})
	}
}

func TestD12RunSystem(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		numSteps int
		run2     bool
		want     string
	}{
		{"test 1", "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", 100, false, "Total Energy: 1940"},
		{"test 2", "<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", 100, true, "Steps to Equilibrium: 2772"},
		{"test 3", "<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", 100, true, "Steps to Equilibrium: 4686774924"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d12 := &D12{}
			err := parseFile(test.input, &challenge{"Test Day 12", "", d12})
			if err == nil {
				got := ""
				if test.run2 {
					got = d12.runSystem2()
				} else {
					got = d12.runSystem(test.numSteps)
				}
				if test.want != got {
					t.Errorf("Wanted %q got %q", test.want, got)
				}
			} else {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestD12Parts(t *testing.T) {
	tests := []challengeTest{}

	for _, test := range tests {
		d12 := &D12{}
		challenge := &challenge{"Test Day 12", "", d12}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}
