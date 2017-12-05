package alg

import "github.com/abates/AdventOfCode/util"

type Coordinate struct {
	*util.Coordinate
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{&util.Coordinate{x, y}}
}

func (c *Coordinate) ID() string {
	return c.String()
}

func (c *Coordinate) Neighbors() []Node {
	nodes := make([]Node, 0)
	c.Coordinate.Neighbors(func(c *util.Coordinate) {
		nodes = append(nodes, &Coordinate{c})
	})
	return nodes
}

func (c *Coordinate) Equal(other *Coordinate) bool { return c.Coordinate.Equal(other.Coordinate) }
