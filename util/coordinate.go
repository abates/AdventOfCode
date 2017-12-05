package util

import (
	"fmt"
)

func ManhattanDistance(c1, c2 *Coordinate) int {
	d := c1.Subtract(c2)
	return Abs(d.X) + Abs(d.Y)
}

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.X, c.Y)
}

func (c *Coordinate) Add(addend *Coordinate) (nextCoordinate *Coordinate) {
	nextCoordinate = &Coordinate{}
	nextCoordinate.X = c.X + addend.X
	nextCoordinate.Y = c.Y + addend.Y
	return nextCoordinate
}

func (c *Coordinate) Subtract(subtrahend *Coordinate) (nextCoordinate *Coordinate) {
	nextCoordinate = &Coordinate{}
	nextCoordinate.X = c.X - subtrahend.X
	nextCoordinate.Y = c.Y - subtrahend.Y
	return nextCoordinate
}

func (c *Coordinate) Equal(other *Coordinate) bool {
	return c.X == other.X && c.Y == other.Y
}

type CoordinateCallback func(*Coordinate)

func (c *Coordinate) Neighbors(callback CoordinateCallback) {
	directions := []Coordinate{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for _, direction := range directions {
		nextCoordinate := c.Add(&direction)
		callback(nextCoordinate)
	}
}

func (c *Coordinate) Diagonals(callback CoordinateCallback) {
	directions := []Coordinate{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}}
	for _, direction := range directions {
		nextCoordinate := c.Add(&direction)
		callback(nextCoordinate)
	}
}
