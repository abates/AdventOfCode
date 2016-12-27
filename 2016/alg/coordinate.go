package alg

import (
	"fmt"
)

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) ID() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func (c *Coordinate) String() string {
	return c.ID()
}

func (c *Coordinate) Add(addend *Coordinate) (nextCoordinate *Coordinate) {
	nextCoordinate = &Coordinate{}
	nextCoordinate.X = c.X + addend.X
	nextCoordinate.Y = c.Y + addend.Y
	return nextCoordinate
}

func (c *Coordinate) Equal(other *Coordinate) bool {
	return c.X == other.X && c.Y == other.Y
}

func (c *Coordinate) Neighbors() []Node {
	coordinates := make([]Node, 0)
	directions := []Coordinate{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for _, direction := range directions {
		nextCoordinate := c.Add(&direction)
		coordinates = append(coordinates, nextCoordinate)
	}
	return coordinates
}
