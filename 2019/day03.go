package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abates/AdventOfCode/coordinate"
)

func init() {
	d3 := &D3{}
	challenges[3] = &challenge{"Day 03", "input/day03.txt", d3}
}

type segment struct {
	*coordinate.Segment
}

type wire struct {
	segments []*segment
}

func (w *wire) distance(c coordinate.Coordinate) int {
	distance := 0
	for _, segment := range w.segments {
		if segment.Bounds(c) {
			distance += int(coordinate.ManhattanDistance(segment.Start, c))
			break
		}
		distance += int(coordinate.ManhattanDistance(segment.Start, segment.End))
	}
	return distance
}

func (w *wire) Intersections(other *wire) []coordinate.Coordinate {
	intersections := []coordinate.Coordinate{}
	for _, s1 := range w.segments {
		for _, s2 := range other.segments {
			if c, i := s1.Intersection(s2.Segment); i {
				if !c.Equal(coordinate.New(0, 0)) {
					intersections = append(intersections, c)
				}
			}
		}
	}
	return intersections
}

func (w *wire) Equal(other *wire) bool {
	if len(w.segments) != len(other.segments) {
		return false
	}

	for i, segment := range w.segments {
		if !segment.Coincident(other.segments[i].Segment) {
			return false
		}
	}
	return true
}

type D3 struct {
	wires []*wire
}

func (d3 *D3) parse(line string) error {
	start := coordinate.New(0, 0)
	end := coordinate.New(0, 0)
	wire := &wire{}
	for _, inst := range strings.Split(line, ",") {
		dir := string([]rune(inst)[0])
		count, err := strconv.Atoi(string([]rune(inst)[1:]))
		if err == nil {
			switch dir {
			case "U":
				end = start.Add(coordinate.New(0, float64(count)))
			case "D":
				end = start.Add(coordinate.New(0, -float64(count)))
			case "L":
				end = start.Add(coordinate.New(-float64(count), 0))
			case "R":
				end = start.Add(coordinate.New(float64(count), 0))
			default:
				return fmt.Errorf("Unknown direction %s", dir)
			}
			wire.segments = append(wire.segments, &segment{coordinate.NewSegment(start, end)})
			start = end
		} else {
			return err
		}
	}
	d3.wires = append(d3.wires, wire)
	return nil
}

func (d3 *D3) closestIntersection() (dist float64, coord coordinate.Coordinate) {
	origin := coordinate.New(0, 0)
	for _, w1 := range d3.wires {
		for _, w2 := range d3.wires {
			if w1 == w2 {
				continue
			}

			for _, c := range w1.Intersections(w2) {
				d := coordinate.ManhattanDistance(origin, c)
				if coord == nil || d < dist {
					dist = d
					coord = c
				}
			}
		}
	}
	return
}

func (d3 *D3) part1() (string, error) {
	d, _ := d3.closestIntersection()
	return fmt.Sprintf("Distance from the central port to the closest intersection is %d", int(d)), nil
}

func (d3 *D3) part2() (string, error) {
	distance := -1
	for _, w1 := range d3.wires {
		for _, w2 := range d3.wires {
			if w1 == w2 {
				continue
			}

			for _, c := range w1.Intersections(w2) {
				d := w1.distance(c) + w2.distance(c)
				if distance < 0 || d < distance {
					distance = d
				}
			}
		}
	}
	return fmt.Sprintf("Distance to closest intersection: %d steps", distance), nil
}
