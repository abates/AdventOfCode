package main

import (
	"testing"
)

func TestExchange(t *testing.T) {
	tests := []struct {
		numElves int
		result   int
		takeFn   TakeFn
	}{
		{5, 3, TakeGiftsLeft},
		{6, 5, TakeGiftsLeft},
		{7, 7, TakeGiftsLeft},
		{8, 1, TakeGiftsLeft},
		{9, 3, TakeGiftsLeft},
		{10, 5, TakeGiftsLeft},
		{11, 7, TakeGiftsLeft},
		{12, 9, TakeGiftsLeft},
		{13, 11, TakeGiftsLeft},
		{14, 13, TakeGiftsLeft},
		{15, 15, TakeGiftsLeft},
		{16, 1, TakeGiftsLeft},
		{5, 2, TakeGiftsAcross},
	}

	for i, test := range tests {
		e := NewWhiteElephantExchange(test.numElves, test.takeFn)
		result := e.Exchange()
		if test.result != result {
			t.Errorf("Test %d expected %d got %d", i, test.result, result)
		}
	}
}
