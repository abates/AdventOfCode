package util

import "testing"

func TestCircularIntList(t *testing.T) {
	list := NewCircularIntList()
	for i := 0; i < 10; i += 2 {
		list.Add(i)
	}

	for i := 0; i < 10; i++ {
		j := list.Next()
		expected := (i % 5) * 2
		if j != expected {
			t.Errorf("test[%d] Expected %d got %d", i, expected, j)
		}

		j = list.Peek()
		expected = ((i + 1) % 5) * 2
		if j != expected {
			t.Errorf("test[%d] Expected %d got %d", i, expected, j)
		}

		j = list.PeekN(3)
		expected = ((i + 3) % 5) * 2
		if j != expected {
			t.Errorf("test[%d] Expected %d got %d", i, expected, j)
		}
	}
}
