package util

import "testing"

func TestSumInt(t *testing.T) {
	sum := SumInt(1, 2, 3, 4, 5)
	if sum != 15 {
		t.Errorf("Expected 15 got %d", sum)
	}
}

func TestSumFloat(t *testing.T) {
	sum := SumFloat(1.0, 2.0, 3.0, 4.0, 5.0)
	if sum != 15.0 {
		t.Errorf("Expected 15 got %d", sum)
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		values []int
		min    int
		max    int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 5},
		{[]int{5, 4, 3, 2, 1}, 1, 5},
		{[]int{}, 0, 0},
		{[]int{-1}, -1, -1},
		{[]int{-5, -4, -3, -2, -1}, -5, -1},
		{[]int{-4, -3, -2, -1, 0}, -4, 0},
	}

	for i, test := range tests {
		max := Max(test.values...)
		if test.max != max {
			t.Errorf("tests[%d] max expected %d got %d", i, test.max, max)
		}

		min := Min(test.values...)
		if test.min != min {
			t.Errorf("tests[%d] min expected %d got %d", i, test.min, max)
		}
	}
}
