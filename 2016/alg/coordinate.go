package alg

import "github.com/abates/AdventOfCode/util"

type Coordinate struct {
	*util.Coordinate
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{util.NewCoordinate(x, y)}
}

func (c *Coordinate) ID() string {
	return c.String()
}

func (c *Coordinate) Neighbors() []Node {
	nodes := make([]Node, 0)
	directions := []*util.Coordinate{util.NewCoordinate(0, -1), util.NewCoordinate(1, 0), util.NewCoordinate(0, 1), util.NewCoordinate(-1, 0)}
	for _, direction := range directions {
		nodes = append(nodes, &Coordinate{c.Add(direction)})
	}

	return nodes
}

func (c *Coordinate) Equal(other *Coordinate) bool { return c.Coordinate.Equal(other.Coordinate) }
