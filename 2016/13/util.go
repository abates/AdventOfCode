package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
)

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.x, c.y)
}

func (c *Coordinate) Add(addend *Coordinate) (nextCoordinate *Coordinate) {
	nextCoordinate = &Coordinate{}
	nextCoordinate.x = c.x + addend.x
	nextCoordinate.y = c.y + addend.y
	return nextCoordinate
}

func (c *Coordinate) Equal(other *Coordinate) bool {
	return c.x == other.x && c.y == other.y
}

type WallDetector func(*Coordinate) bool

func MagicDetector(magicNumber int) WallDetector {
	return func(coordinate *Coordinate) bool {
		value := coordinate.x*coordinate.x + 3*coordinate.x + 2*coordinate.x*coordinate.y + coordinate.y + coordinate.y*coordinate.y + magicNumber
		count := 0
		for value > 0 {
			if value&0x01 == 1 {
				count++
			}
			value = value >> 1
		}

		return count%2 == 0
	}
}

func Draw(width, height int, detector WallDetector) string {
	writer := util.StringWriter{}
	writer.Write("  ")
	for x := 0; x < width; x++ {
		writer.Writef("%d", x)
	}
	writer.Write("\n")

	for y := 0; y < height; y++ {
		writer.Writef("%d ", y)
		for x := 0; x < width; x++ {
			if detector(&Coordinate{x, y}) {
				writer.Write(".")
			} else {
				writer.Write("#")
			}
		}
		writer.Write("\n")
	}

	return writer.String()
}

type Walker struct {
	position    *Coordinate
	isOpenSpace WallDetector
	visited     map[int]map[int]bool
}

func NewWalker(startX, startY int, wallDetector WallDetector) *Walker {
	return &Walker{
		position:    &Coordinate{startX, startY},
		isOpenSpace: wallDetector,
		visited:     make(map[int]map[int]bool),
	}
}

func NewMagicWalker(startX, startY, magicNumber int) *Walker {
	return NewWalker(startX, startY, MagicDetector(magicNumber))
}

func (w *Walker) CanWalk(coordinate *Coordinate) bool {
	if coordinate.x < 0 || coordinate.y < 0 {
		return false
	}
	return w.isOpenSpace(coordinate)
}

func (w *Walker) Next() []*Coordinate {
	coordinates := make([]*Coordinate, 0)
	possibleCoordinates := []Coordinate{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	for _, delta := range possibleCoordinates {
		nextCoordinate := w.position.Add(&delta)
		if _, found := w.visited[nextCoordinate.x][nextCoordinate.y]; !found && w.CanWalk(nextCoordinate) {
			coordinates = append(coordinates, nextCoordinate)
		}
	}
	return coordinates
}

var ErrNotFound = fmt.Errorf("Path did not end at destination")

func (w *Walker) Walk(destination Coordinate) int {
	return w.walk(destination, 0)
}

func (w *Walker) WalkMax(destination Coordinate, maxSteps int) int {
	w.walk(destination, maxSteps)
	count := 0
	for _, x := range w.visited {
		for range x {
			count++
		}
	}
	return count
}

type PathCoordinate struct {
	*Coordinate
	depth int
}

type CoordinateQueue struct {
	queue []*PathCoordinate
}

func (c *CoordinateQueue) Push(coordinates ...*PathCoordinate) {
	c.queue = append(c.queue, coordinates...)
}

func (c *CoordinateQueue) Shift() (coordinate *PathCoordinate) {
	if len(c.queue) > 0 {
		coordinate = c.queue[0]
		c.queue = c.queue[1:]
	}
	return
}

func (c *CoordinateQueue) Len() int {
	return len(c.queue)
}

func (w *Walker) walk(destination Coordinate, maxSteps int) int {
	q := CoordinateQueue{
		queue: make([]*PathCoordinate, 0),
	}

	q.Push(&PathCoordinate{w.position, 0})

	depth := 0
	for q.Len() > 0 {
		nextCoordinate := q.Shift()
		w.position = nextCoordinate.Coordinate
		depth = nextCoordinate.depth

		if maxSteps > 0 && depth > maxSteps {
			break
		}

		if _, found := w.visited[w.position.x]; !found {
			w.visited[w.position.x] = make(map[int]bool)
		}
		w.visited[w.position.x][w.position.y] = true

		if w.position.Equal(&destination) {
			break
		}

		for _, coordinate := range w.Next() {
			q.Push(&PathCoordinate{coordinate, depth + 1})
		}
	}

	return depth
}
