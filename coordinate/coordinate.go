package coordinate

import (
	"fmt"
	"math"
	"strings"
)

func EuclideanDistance(c1, c2 Coordinate) float64 {
	d := c1.Subtract(c2)
	d = d.Multiply(d)
	distance := float64(0)
	for i := 0; i < d.Cardinality(); i++ {
		distance += d.Get(i)
	}
	return math.Sqrt(float64(distance))
}

func ManhattanDistance(c1, c2 Coordinate) float64 {
	d := c1.Subtract(c2)
	distance := float64(0)
	for i := 0; i < d.Cardinality(); i++ {
		distance += math.Abs(d.Get(i))
	}
	return distance
}

type Coordinate interface {
	Get(i int) float64
	Cardinality() int
	Add(addend Coordinate) Coordinate
	Subtract(subtrahend Coordinate) Coordinate
	Multiply(multiplier Coordinate) Coordinate
	Equal(other Coordinate) bool
}

type coordinate struct {
	coordinates []float64
}

func New(coordinates ...float64) Coordinate {
	return &coordinate{coordinates}
}

func (c *coordinate) Cardinality() int { return len(c.coordinates) }

func (c *coordinate) Get(i int) float64 { return c.coordinates[i] }

func (c *coordinate) String() string {
	str := make([]string, len(c.coordinates))
	for i, coordinate := range c.coordinates {
		str[i] = fmt.Sprintf("%f", coordinate)
	}
	return fmt.Sprintf("(%s)", strings.Join(str, ","))
}

func (c *coordinate) compute(other Coordinate, computer func(lhs, rhs float64) float64) Coordinate {
	nextCoordinate := &coordinate{
		coordinates: make([]float64, len(c.coordinates)),
	}
	for i, coordinate := range c.coordinates {
		nextCoordinate.coordinates[i] = computer(coordinate, other.Get(i))
	}
	return nextCoordinate
}

func (c *coordinate) Add(addend Coordinate) Coordinate {
	return c.compute(addend, func(lhs, rhs float64) float64 { return lhs + rhs })
}

func (c *coordinate) Subtract(subtrahend Coordinate) Coordinate {
	return c.compute(subtrahend, func(lhs, rhs float64) float64 { return lhs - rhs })
}

func (c *coordinate) Multiply(multiplier Coordinate) Coordinate {
	return c.compute(multiplier, func(lhs, rhs float64) float64 { return lhs * rhs })
}

func (c *coordinate) Equal(other Coordinate) bool {
	if c == other {
		return true
	}

	if c == nil || other == nil {
		return false
	}

	for i, coordinate := range c.coordinates {
		if coordinate != other.Get(i) {
			return false
		}
	}
	return true
}

func NewSegment(start, end Coordinate) *Segment {
	seg := &Segment{
		Start: start,
		End:   end,
		s:     end.Subtract(start),
	}

	return seg
}

type Segment struct {
	Start Coordinate
	End   Coordinate
	s     Coordinate
}

func (seg *Segment) String() string {
	return fmt.Sprintf("%v -> %v", seg.Start, seg.End)
}

func (seg *Segment) Coincident(other *Segment) bool {
	return seg.Start.Equal(other.Start) && seg.End.Equal(other.End)
}

func (seg *Segment) Intersection(other *Segment) (Coordinate, bool) {
	d := float64((other.s.Get(1)*seg.s.Get(0) - other.s.Get(0)*seg.s.Get(1)))

	m1 := float64((other.s.Get(0) * (seg.Start.Get(1) - other.Start.Get(1))) - (other.s.Get(1) * (seg.Start.Get(0) - other.Start.Get(0))))
	m2 := float64(seg.s.Get(0)*(seg.Start.Get(1)-other.Start.Get(1)) - (seg.s.Get(1) * (seg.Start.Get(0) - other.Start.Get(0))))

	if d == 0 {
		return nil, false
	}

	m1 = m1 / d
	m2 = m2 / d

	// find the intersection of the *lines* (not segment)
	intersection := &coordinate{
		coordinates: make([]float64, seg.Start.Cardinality()),
	}

	for i := 0; i < seg.Start.Cardinality(); i++ {
		a := seg.Start.Get(i) + seg.s.Get(i)*m1
		b := other.Start.Get(i) + other.s.Get(i)*m2
		if a != b {
			return nil, false
		}
		intersection.coordinates[i] = a
	}

	// determine if the intersection is within both lines
	if seg.Contains(intersection) && other.Contains(intersection) {
		return intersection, true
	}
	return nil, false
}

func (seg *Segment) Contains(point Coordinate) bool {
	for i := 0; i < seg.Start.Cardinality(); i++ {
		start := seg.Start.Get(i)
		end := seg.End.Get(i)
		if end < start {
			t := end
			end = start
			start = t
		}

		if point.Get(i) < start || end < point.Get(i) {
			return false
		}
	}
	return true
}

func (seg *Segment) Length() int {
	return int(EuclideanDistance(seg.Start, seg.End))
}
