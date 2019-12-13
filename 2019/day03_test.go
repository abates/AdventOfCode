package main

import (
	"strings"
	"testing"

	"github.com/abates/AdventOfCode/coordinate"
)

func TestD3Parse(t *testing.T) {
	cs := func(x1, y1, x2, y2 float64) *segment {
		return &segment{coordinate.NewSegment(coordinate.New(x1, y1), coordinate.New(x2, y2))}
	}

	tests := []struct {
		name  string
		input string
		want  *wire
	}{
		{"test 1", "R8,U5,L5,D3", &wire{[]*segment{cs(0, 0, 8, 0), cs(8, 0, 8, 5), cs(8, 5, 3, 5), cs(3, 5, 3, 2)}}},
	}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			d3 := &D3{}
			err := d3.parse(test.input)
			if err == nil {
				if len(d3.wires) == 1 {
					if !test.want.Equal(d3.wires[0]) {
						t.Errorf("Wanted wire %v got %v", test.want, d3.wires[0])
					}
				} else {
					t.Errorf("Expected 1 wire, got %d", len(d3.wires))
				}
			} else {
				t.Errorf("Unexpected Error: %v", err)
			}
		})
	}
}

func TestD3WireIntersections(t *testing.T) {
	nc := func(x, y float64) coordinate.Coordinate {
		return coordinate.New(x, y)
	}

	tests := []struct {
		name  string
		wire1 string
		wire2 string
		want  []coordinate.Coordinate
	}{
		{"test 1", "R8,U5,L5,D3", "U7,R6,D4,L4", []coordinate.Coordinate{nc(3, 3), nc(6, 5)}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d3 := &D3{}
			for _, line := range []string{test.wire1, test.wire2} {
				err := d3.parse(line)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if len(d3.wires) == 2 {
				got := d3.wires[0].Intersections(d3.wires[1])
				for _, g1 := range got {
					found := false
					for _, w1 := range test.want {
						if w1.Equal(g1) {
							found = true
							break
						}
					}

					if !found {
						t.Errorf("Wanted all of %v in %v", test.want, got)
					}
				}
			} else {
				t.Errorf("Wanted 2 wires, got %d", len(d3.wires))
			}
		})
	}
}

func TestD3ClosestInstersection(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantDist  float64
		wantCoord coordinate.Coordinate
	}{
		{"test 1", "R8,U5,L5,D3\nU7,R6,D4,L4", 6, coordinate.New(3, 3)},
		{"test 2", "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159, coordinate.New(155, 4)},
		{"test 2", "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135, coordinate.New(124, 11)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d3 := &D3{}
			for _, line := range strings.Split(test.input, "\n") {
				err := d3.parse(line)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			gotDist, gotCoord := d3.closestIntersection()
			if test.wantDist == gotDist {
				if !test.wantCoord.Equal(gotCoord) {
					t.Errorf("Wanted intersection %v got %v", test.wantCoord, gotCoord)
				}
			} else {
				t.Errorf("Wanted distance %f got %f", test.wantDist, gotDist)
			}
		})
	}
}

func TestD3Parts(t *testing.T) {
	tests := []challengeTest{
		{"test 1", "R8,U5,L5,D3\nU7,R6,D4,L4", "Distance from the central port to the closest intersection is 6", ""},
		{"test 2", "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", "Distance from the central port to the closest intersection is 159", "Distance to closest intersection: 610 steps"},
		{"test 2", "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", "Distance from the central port to the closest intersection is 135", "Distance to closest intersection: 410 steps"},
	}

	for _, test := range tests {
		d3 := &D3{}
		challenge := &challenge{"Test Day 03", "", d3.parse, nil, d3.part1, d3.part2}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}
