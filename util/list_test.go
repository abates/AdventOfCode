package util

import (
	"reflect"
	"testing"
)

func TestCircularIntList(t *testing.T) {
	list := CircularIntList{}
	expected := make([]int, 0)
	for i := 0; i < 10; i += 2 {
		expected = append(expected, i)
		list.Add(i)
	}

	values := make([]int, 0)
	list.Iterate(func(i int) {
		values = append(values, i)
	})

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("test expected %-v got %-v", expected, values)
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
