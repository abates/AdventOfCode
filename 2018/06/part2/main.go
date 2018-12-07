package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Coordinate struct {
	x int
	y int
}

func abs(input int) int {
	if input < 0 {
		return -1 * input
	}
	return input
}

func (c *Coordinate) Distance(x, y int) int {
	return abs(x-c.x) + abs(y-c.y)
}

func (c *Coordinate) UnmarshalText(text []byte) error {
	str := string(text)
	_, err := fmt.Sscanf(str, "%d, %d", &c.x, &c.y)
	return err
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
		err = grid.AddString(line)
		if err != nil {
			panic("Error: " + err.Error())
		}
	}

	area := 0
	for x := grid.minX; x <= grid.maxX; x++ {
		for y := grid.minY; y < grid.maxY; y++ {
			sum := 0
			for _, coordinate := range grid.coordinates {
				sum += coordinate.Distance(x, y)
			}
			if sum < 10000 {
				area++
			}
		}
	}
	fmt.Printf("Area: %d\n", area)
}
