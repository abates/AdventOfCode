package util

import (
	"fmt"
	"strings"
)

func ManhattanDistance(c1, c2 *Coordinate) int {
	d := c1.Subtract(c2)
	distance := 0
	for _, coordinate := range d.Coordinates {
		distance += Abs(coordinate)
	}
	return distance
}

type Coordinate struct {
	Coordinates []int
}

func NewCoordinate(coordinates ...int) *Coordinate {
	return &Coordinate{coordinates}
}

func (c *Coordinate) String() string {
	str := make([]string, len(c.Coordinates))
	for i, coordinate := range c.Coordinates {
		str[i] = fmt.Sprintf("%d", coordinate)
	}
	return fmt.Sprintf("(%s)", strings.Join(str, ","))
}

func (c *Coordinate) Add(addend *Coordinate) (nextCoordinate *Coordinate) {
	nextCoordinate = &Coordinate{
		Coordinates: make([]int, len(c.Coordinates)),
	}
	for i, coordinate := range c.Coordinates {
		nextCoordinate.Coordinates[i] = coordinate + addend.Coordinates[i]
	}
	return nextCoordinate
}

func (c *Coordinate) Subtract(subtrahend *Coordinate) (nextCoordinate *Coordinate) {
	nextCoordinate = &Coordinate{
		Coordinates: make([]int, len(c.Coordinates)),
	}
	for i, coordinate := range c.Coordinates {
		nextCoordinate.Coordinates[i] = coordinate - subtrahend.Coordinates[i]
	}
	return nextCoordinate
}

func (c *Coordinate) Equal(other *Coordinate) bool {
	for i, coordinate := range c.Coordinates {
		if coordinate != other.Coordinates[i] {
			return false
		}
	}
	return true
}

/*type CoordinateCallback func(*Coordinate)

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
}*/
