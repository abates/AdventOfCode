package main

import (
	"testing"
)

func TestSetIntersection(t *testing.T) {
	tests := []struct {
		s1       []string
		s2       []string
		expected string
	}{
		{[]string{"a", "b"}, []string{"b", "c"}, "[b]"},
		{[]string{"c", "b"}, []string{"b", "c"}, "[b, c]"},
		{[]string{"a", "b", "c", "d"}, []string{"b"}, "[b]"},
		{[]string{"b"}, []string{"a", "b", "c", "d"}, "[b]"},
	}

	for i, test := range tests {
		s1 := make(Set)
		s2 := make(Set)
		for _, s := range test.s1 {
			s1.Add(s)
		}

		for _, s := range test.s2 {
			s2.Add(s)
		}
		intersection := s1.Intersection(s2)
		if intersection.String() != test.expected {
			t.Errorf("tests[%d] expected %q got %q", i, test.expected, intersection.String())
		}
	}
}
