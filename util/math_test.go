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
