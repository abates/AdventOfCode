package coordinate

import (
	"fmt"
	"testing"
)

func TestEuclideanDistance(t *testing.T) {
	tests := []struct {
		c1   []float64
		c2   []float64
		want float64
	}{
		{[]float64{2, -1}, []float64{-2, 2}, 5},
	}

	for _, test := range tests {
		c1 := New(test.c1...)
		c2 := New(test.c2...)
		t.Run(fmt.Sprintf("Euclidean distance between %s and %s", c1, c2), func(t *testing.T) {
			got := EuclideanDistance(c1, c2)
			if test.want != got {
				t.Errorf("Wanted %v got %v", test.want, got)
			}
		})
	}
}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		x1       float64
		y1       float64
		x2       float64
		y2       float64
		expected float64
	}{
		{0, 0, 1, 1, 2},
		{0, 0, 2, 2, 4},
		{-1, -1, 1, 1, 4},
		{1, 1, 1, 1, 0},
	}

	for i, test := range tests {
		start := New(test.x1, test.y1)
		end := New(test.x2, test.y2)

		d := ManhattanDistance(start, end)
		if d != test.expected {
			t.Errorf("tests[%d] expected %f got %f", i, test.expected, d)
		}

		d = ManhattanDistance(end, start)
		if d != test.expected {
			t.Errorf("tests[%d] expected %f got %f", i, test.expected, d)
		}
	}
}

func TestAddSubtractCoordinates(t *testing.T) {
	tests := []struct {
		x     float64
		y     float64
		diffX float64
		diffY float64
		addX  float64
		addY  float64
		subX  float64
		subY  float64
	}{
		{1, 1, 0, -1, 1, 0, 1, 2},
		{1, 1, 0, -1, 1, 0, 1, 2},
		{1, 1, 0, -1, 1, 0, 1, 2},
		{1, 1, -1, 0, 0, 1, 2, 1},
	}

	for i, test := range tests {
		coord := New(test.x, test.y)
		diffX := New(test.diffX, test.diffY)

		result := coord.Add(diffX)
		expected := New(test.addX, test.addY)

		wantStr := fmt.Sprintf("(%f,%f)", test.x, test.y)
		gotStr := fmt.Sprintf("%s", coord)
		if wantStr != gotStr {
			t.Errorf("tests[%d] expected %s got %s", i, wantStr, gotStr)
		}

		if !result.Equal(expected) {
			t.Errorf("tests[%d] Add expected %s got %s", i, expected, result)
		}

		result = coord.Subtract(diffX)
		expected = New(test.subX, test.subY)

		if !result.Equal(expected) {
			t.Errorf("tests[%d] Subtract expected %s got %s", i, expected, result)
		}
	}
}

func TestSegmentCoincident(t *testing.T) {
	tests := []struct {
		name string
		s1   *Segment
		s2   *Segment
		want bool
	}{
		{"true", NewSegment(New(0, 0), New(0, 1)), NewSegment(New(0, 0), New(0, 1)), true},
		{"false", NewSegment(New(0, 2), New(0, 1)), NewSegment(New(0, 0), New(0, 1)), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, s := range [][]*Segment{{test.s1, test.s2}, {test.s2, test.s1}} {
				got := s[0].Coincident(s[1])
				if test.want != got {
					t.Errorf("Wanted (%v).Coincident(%v) == %v got %v", s[0], s[1], test.want, got)
				}
			}
		})
	}
}

func TestSegmentIntersection(t *testing.T) {
	tests := []struct {
		name             string
		l1               *Segment
		l2               *Segment
		wantIntersection Coordinate
		wantIntersect    bool
	}{
		{"does not intersect", NewSegment(New(5, 2, -1), New(6, 0, -4)), NewSegment(New(2, 0, 4), New(3, 2, 3)), nil, false},
		{"intersect", NewSegment(New(3, 6, 5), New(15, -18, -31)), NewSegment(New(1, -2, 5), New(12, 20, -6)), New(4, 4, 2), true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotIntersection, gotIntersect := test.l1.Intersection(test.l2)

			if gotIntersect == test.wantIntersect {
				if gotIntersect && !gotIntersection.Equal(test.wantIntersection) {
					t.Errorf("Wanted intersection %v got %v", test.wantIntersection, gotIntersection)
				}
			} else {
				t.Errorf("Wanted intersect to be %v got %v", test.wantIntersect, gotIntersect)
			}
		})
	}
}
