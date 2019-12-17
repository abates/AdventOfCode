package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/abates/AdventOfCode/coordinate"
)

func init() {
	d10 := &D10{}
	challenges[10] = &challenge{"Day 10", "input/day10.txt", d10}
}

type clockwise struct {
	asteroids
	center coordinate.Coordinate
}

func (c clockwise) Less(i, j int) bool {
	a := c.asteroids[i]
	b := c.asteroids[j]

	return math.Atan2(a.Get(0)-c.center.Get(0), a.Get(1)-c.center.Get(1)) > math.Atan2(b.Get(0)-c.center.Get(0), b.Get(1)-c.center.Get(1))
}

type asteroids []coordinate.Coordinate

func (a asteroids) Len() int      { return len(a) }
func (a asteroids) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a *asteroids) append(coordinates ...coordinate.Coordinate) {
	for _, c := range coordinates {
		(*a) = append(*a, c)
	}
}

func (a asteroids) best() (center coordinate.Coordinate, visible asteroids) {
	center = a[0]
	visible = a.visibleFrom(center)
	for _, a1 := range a[1:] {
		v := a.visibleFrom(a1)
		if len(v) > len(visible) {
			visible = v
			center = a1
		}
	}
	sort.Sort(&clockwise{visible, center})
	return
}

func (a asteroids) visibleFrom(a1 coordinate.Coordinate) asteroids {
	index := make(map[string]*coordinate.Segment)
	for _, a2 := range a {
		if a1 == a2 {
			continue
		}

		line := coordinate.NewSegment(a1, a2)
		v := fmt.Sprintf("%v", coordinate.UnitVector(line.Direction))
		if s, found := index[v]; found {
			if line.Bounds(s.End) {
				// new coordinate is blocked by s
				continue
			} else {
				// old coordinate is blocked by s
				index[v] = line
			}
		} else {
			index[v] = line
		}
	}

	visible := asteroids{}
	for _, s := range index {
		visible = append(visible, s.End)
	}
	return visible
}

type D10 struct {
	rows      int
	asteroids asteroids
}

func (d10 *D10) parse(line string) error {
	y := float64(d10.rows)
	for col, s := range strings.Split(line, "") {
		x := float64(col)
		if s == "#" {
			d10.asteroids.append(coordinate.New(x, y))
		} else if s != "." {
			return fmt.Errorf("Unknown character %q", s)
		}
	}
	d10.rows++
	return nil
}

func (d10 *D10) part1() (string, error) {
	center, visible := d10.asteroids.best()
	return fmt.Sprintf("Best Location: (%d, %d): %d asteroids visible", int(center.Get(0)), int(center.Get(1)), len(visible)), nil
}

func (d10 *D10) part2() (string, error) {
	center, _ := d10.asteroids.best()

	am := make(map[string]coordinate.Coordinate)
	for _, c := range d10.asteroids {
		am[fmt.Sprintf("%v", c)] = c
	}

	for i := 0; i <= 200 && len(am) > 1; {
		a := asteroids{}
		for _, c := range am {
			a = append(a, c)
		}
		visible := a.visibleFrom(center)
		sort.Sort(&clockwise{visible, center})
		for _, v := range visible {
			delete(am, fmt.Sprintf("%v", v))
			i++
			if i == 200 {
				return fmt.Sprintf("Value for 200th asteroid destroyed: %v", v.Get(0)*100+v.Get(1)), nil
			}
		}
	}

	return "unknown", fmt.Errorf("Shouldn't get here")
}
