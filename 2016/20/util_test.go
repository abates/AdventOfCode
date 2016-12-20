package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		r1     *Range
		r2     *Range
		result *Range
	}{
		{&Range{0, 5}, &Range{4, 10}, &Range{0, 10}},
	}

	for i, test := range tests {
		test.r1.Merge(test.r2)
		if test.r1.low != test.result.low || test.r1.high != test.result.high {
			t.Errorf("Test %d expected %s got %s", i, test.result, test.r1)
		}
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		r1     *Range
		r2     *Range
		result int
	}{
		{&Range{4, 10}, &Range{0, 5}, 0},
		{&Range{4, 10}, &Range{0, 3}, -1},
		{&Range{4, 10}, &Range{11, 15}, 1},
		{&Range{4, 10}, &Range{5, 9}, 0},
	}

	for i, test := range tests {
		result := test.r1.Compare(test.r2)
		if result != test.result {
			t.Errorf("Test %d expected %d got %d", i, test.result, result)
		}
	}
}

func TestAdjacent(t *testing.T) {
	tests := []struct {
		r1     *Range
		r2     *Range
		result bool
	}{
		{&Range{6, 10}, &Range{0, 5}, true},
		{&Range{4, 10}, &Range{12, 15}, false},
	}

	for i, test := range tests {
		result := test.r1.Adjacent(test.r2)
		if result != test.result {
			t.Errorf("Test %d expected %d got %d", i, test.result, result)
		}
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		r1     *Range
		result int
	}{
		{&Range{6, 10}, 5},
		{&Range{0, 2}, 3},
	}

	for i, test := range tests {
		result := test.r1.Count()
		if result != test.result {
			t.Errorf("Test %d expected %d got %d", i, test.result, result)
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		ranges []*Range
		result []*Range
	}{
		{[]*Range{&Range{8, 10}, &Range{4, 6}, &Range{0, 2}}, []*Range{&Range{0, 2}, &Range{4, 6}, &Range{8, 10}}},
		{[]*Range{&Range{8, 10}, &Range{4, 6}, &Range{0, 2}, &Range{7, 15}}, []*Range{&Range{0, 2}, &Range{4, 15}}},
		{[]*Range{&Range{8, 10}, &Range{4, 6}, &Range{0, 2}, &Range{11, 15}}, []*Range{&Range{0, 2}, &Range{4, 6}, &Range{8, 15}}},
		{[]*Range{&Range{8, 10}, &Range{4, 6}, &Range{0, 2}, &Range{1, 3}}, []*Range{&Range{0, 6}, &Range{8, 10}}},
		{[]*Range{&Range{20, 25}, &Range{10, 15}, &Range{0, 5}, &Range{6, 9}}, []*Range{&Range{0, 15}, &Range{20, 25}}},
	}

	for i, test := range tests {
		ranges := &Ranges{}
		for _, rng := range test.ranges {
			ranges.Insert(rng)
		}

		if !reflect.DeepEqual(test.result, ranges.ranges) {
			t.Errorf("Test %d expected %v got %v", i, test.result, ranges.ranges)
		}
	}
}

func TestLowestException(t *testing.T) {
	tests := []struct {
		ranges []*Range
		result int
	}{
		{[]*Range{&Range{0, 2}, &Range{4, 6}, &Range{8, 10}}, 3},
		{[]*Range{&Range{0, 6}, &Range{8, 15}}, 7},
	}

	for i, test := range tests {
		ranges := &Ranges{test.ranges}
		result := ranges.LowestException()
		if result != test.result {
			t.Errorf("Test %d expected %v got %v", i, test.result, result)
		}
	}
}
