package alg

import "github.com/abates/AdventOfCode/coordinate"

type Coordinate struct {
	*coordinate.Coordinate
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{coordinate.New(x, y)}
}

func (c *Coordinate) ID() string {
	return c.String()
}

func (c *Coordinate) Neighbors() []Node {
	nodes := make([]Node, 0)
	directions := []*coordinate.Coordinate{coordinate.New(0, -1), coordinate.New(1, 0), coordinate.New(0, 1), coordinate.New(-1, 0)}
	for _, direction := range directions {
		nodes = append(nodes, &Coordinate{c.Add(direction)})
	}

	return nodes
}

func (c *Coordinate) Equal(other *Coordinate) bool { return c.Coordinate.Equal(other.Coordinate) }
