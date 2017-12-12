package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/abates/AdventOfCode/util"
)

type Direction *util.Coordinate

var (
	North     = Direction(util.NewCoordinate(0, 1, -1))
	NorthEast = Direction(util.NewCoordinate(1, 0, -1))
	SouthEast = Direction(util.NewCoordinate(1, -1, 0))
	South     = Direction(util.NewCoordinate(0, -1, 1))
	SouthWest = Direction(util.NewCoordinate(-1, 0, 1))
	NorthWest = Direction(util.NewCoordinate(-1, 1, 0))

	Directions = map[string]Direction{
		"n":  North,
		"ne": NorthEast,
		"se": SouthEast,
		"s":  South,
		"sw": SouthWest,
		"nw": NorthWest,
	}
)

type Mover struct {
	coordinate *util.Coordinate
}

func (m *Mover) Move(direction Direction) {
	m.coordinate = m.coordinate.Add(direction)
}

func distance(mover *Mover) int {
	distance := 0
	for _, coordinate := range mover.coordinate.Coordinates {
		distance += util.Abs(coordinate)
	}
	return distance / 2
}

func positionChild(input string) (*Mover, int) {
	m := &Mover{util.NewCoordinate(0, 0, 0)}
	maxDistance := 0
	for _, direction := range strings.Split(input, ",") {
		direction = strings.TrimSpace(direction)
		if direction == "" {
			continue
		}
		m.Move(Directions[direction])
		d := distance(m)
		if maxDistance < d {
			maxDistance = d
		}
	}
	return m, maxDistance
}

func main() {
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	m, max := positionChild(string(b))
	fmt.Printf("Child is now at %s\n", m.coordinate.String())
	fmt.Printf("Max distance is %d\n", max)
	d := distance(m)
	fmt.Printf("Distance is %d\n", d)
}
