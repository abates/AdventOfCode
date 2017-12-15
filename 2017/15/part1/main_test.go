package main

import "testing"

func TestIncrement(t *testing.T) {
	tests := []struct {
		input     int
		factor    int
		multiple  int
		increment int
		expected  int
	}{
		{65, 16807, 1, 1, 1092455},
		{65, 16807, 1, 2, 1181022009},
		{65, 16807, 1, 3, 245556042},
		{65, 16807, 1, 4, 1744312007},
		{65, 16807, 1, 5, 1352636452},

		{65, 16807, 4, 1, 1352636452},
		{65, 16807, 4, 2, 1992081072},
		{65, 16807, 4, 3, 530830436},
		{65, 16807, 4, 4, 1980017072},
		{65, 16807, 4, 5, 740335192},

		{8921, 48271, 1, 1, 430625591},
		{8921, 48271, 1, 2, 1233683848},
		{8921, 48271, 1, 3, 1431495498},
		{8921, 48271, 1, 4, 137874439},
		{8921, 48271, 1, 5, 285222916},

		{8921, 48271, 8, 1, 1233683848},
		{8921, 48271, 8, 2, 862516352},
		{8921, 48271, 8, 3, 1159784568},
		{8921, 48271, 8, 4, 1616057672},
		{8921, 48271, 8, 5, 412269392},
	}

	for i, test := range tests {
		gen := NewGenerator(test.input, test.factor, test.multiple)
		for j := 0; j < test.increment-1; j++ {
			gen.Next()
		}
		value := gen.Next()
		if test.expected != value {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, value)
		}
	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		value1   int
		value2   int
		expected bool
	}{
		{1092455, 430625591, false},
		{1181022009, 1233683848, false},
		{245556042, 1431495498, true},
		{1744312007, 137874439, false},
		{1352636452, 285222916, false},
	}

	for i, test := range tests {
		if test.expected != Match(test.value1, test.value2) {
			t.Errorf("tests[%d] expected %v got %v", i, test.expected, Match(test.value1, test.value2))
		}
	}
}

func TestCountMatches(t *testing.T) {
	tests := []struct {
		genASeed     int
		genAFactor   int
		genAMultiple int
		genBSeed     int
		genBFactor   int
		genBMultiple int
		limit        int
		expected     int
	}{
		{65, 16807, 1, 8921, 48271, 1, 5, 1},
		{65, 16807, 4, 8921, 48271, 8, 1057, 1},
	}

	for i, test := range tests {
		genA := NewGenerator(test.genASeed, test.genAFactor, test.genAMultiple)
		genB := NewGenerator(test.genBSeed, test.genBFactor, test.genBMultiple)

		count := CountMatches(genA, genB, test.limit)
		if test.expected != count {
			t.Errorf("tests[%d] expected %d got %d", i, test.expected, count)
		}
	}
}
