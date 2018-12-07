package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Coordinate struct {
	x          int
	y          int
	area       int
	isInfinite bool
	distance   int
}

func (c *Coordinate) UnmarshalText(text []byte) error {
	str := string(text)
	_, err := fmt.Sscanf(str, "%d, %d", &c.x, &c.y)
	return err
}

type Distance struct {
	owner    *Coordinate
	distance int
}

type Grid struct {
	coordinates []*Coordinate
	minX        int
	maxX        int
	minY        int
	maxY        int
}

func (g *Grid) AddString(text []byte) error {
	c := &Coordinate{}
	err := c.UnmarshalText(text)
	if err == nil {
		g.Add(c)
	}
	return err
}

func (g *Grid) Add(c *Coordinate) {
	// First updates the bounds of the box that
	// contains the coordinates
	if len(g.coordinates) == 0 {
		g.minX = c.x
		g.maxX = c.x
		g.minY = c.y
		g.maxY = c.y
	}

	if c.x < g.minX {
		g.minX = c.x
	}
	if c.x > g.maxX {
		g.maxX = c.x
	}
	if c.y < g.minY {
		g.minY = c.y
	}
	if c.y > g.maxY {
		g.maxY = c.y
	}
	g.coordinates = append(g.coordinates, c)
}

type Search struct {
	grid    *Grid
	owner   *Coordinate
	results map[string]struct{}
}

func (s *Search) Search() []string {
	s.search(s.owner.x, s.owner.y, s.owner.distance)
	results := []string{}
	for k := range s.results {
		results = append(results, k)
	}
	return results
}

func (s *Search) add(x, y int) {
	if s.results == nil {
		s.results = make(map[string]struct{})
	}
	x = x + s.owner.x
	y = y + s.owner.y
	// don't count coordinates outside the bounding box
	if x >= s.grid.minX && x <= s.grid.maxX && y >= s.grid.minY && y <= s.grid.maxY {
		// add unique coordinates
		s.results[fmt.Sprintf("(%d, %d)", x, y)] = struct{}{}
	}
}

// Find all the x,y values within <<distance>>
func (s *Search) search(x, y, distance int) {
	for x := 0; x <= distance; x++ {
		s.add(x, distance-x)
		s.add(x, -1*(distance-x))
	}

	for x := -1 * distance; x <= 0; x++ {
		s.add(x, distance+x)
		s.add(x, -1*(distance+x))
	}
}

func (g *Grid) Search() *Coordinate {
	index := make(map[string]*Distance)
	queue := make([]*Coordinate, len(g.coordinates))
	coordinates := make([]*Coordinate, len(g.coordinates))
	for i, c := range g.coordinates {
		newC := &Coordinate{}
		*newC = *c
		queue[i] = newC
		coordinates[i] = newC
		index[fmt.Sprintf("(%d, %d)", newC.x, newC.y)] = &Distance{owner: newC, distance: 0}

		if newC.x == g.minX || newC.x == g.maxX || newC.y == g.minY || newC.y == g.maxY {
			newC.isInfinite = true
		}
	}

	// Update the areas and distances
	update := func(coordinate *Coordinate, distance int, keys []string) []string {
		retained := []string{}
		for _, key := range keys {
			if d, found := index[key]; found {
				if distance < d.distance {
					d.owner.area--
					d.owner = coordinate
					d.distance = distance
					coordinate.area++
					retained = append(retained, key)
				} else if distance == d.distance && d.owner != nil {
					d.owner.area--
					d.owner = nil
				}
			} else {
				index[key] = &Distance{
					owner:    coordinate,
					distance: distance,
				}
				coordinate.area++
				retained = append(retained, key)
			}
		}
		return retained
	}

	for len(queue) > 0 {
		coordinate := queue[0]
		coordinate.distance++
		queue = queue[1:]

		s := &Search{grid: g, owner: coordinate}
		candidates := s.Search()
		updates := update(coordinate, coordinate.distance, candidates)
		if len(updates) > 0 {
			queue = append(queue, coordinate)
		}
	}

	maxArea := coordinates[0].area
	maxCoordinate := coordinates[0]
	for _, coordinate := range coordinates[1:] {
		if !coordinate.isInfinite && coordinate.area > maxArea {
			maxArea = coordinate.area
			maxCoordinate = coordinate
		}
	}
	return maxCoordinate
}

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Error: " + err.Error())
	}

	grid := &Grid{}
	for _, line := range bytes.Split(input, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		err := grid.AddString(line)
		if err != nil {
			panic("Error: " + err.Error())
		}
	}

	c := grid.Search()
	// area doesn't include the coordinate itself
	fmt.Printf("Area: %d\n", c.area+1)
}
