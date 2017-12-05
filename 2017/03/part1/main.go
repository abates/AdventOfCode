// I hate this solution... :-/
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abates/AdventOfCode/util"
)

var values = make(map[string]int)

type Direction string

const (
	NORTH = Direction("north")
	SOUTH = Direction("south")
	EAST  = Direction("east")
	WEST  = Direction("west")
)

type Mover struct {
	coordinate *util.Coordinate
	id         int
}

func (m *Mover) Move(direction Direction) {
	m.id++
	switch direction {
	case NORTH:
		m.coordinate = m.coordinate.Add(&util.Coordinate{0, 1})
	case SOUTH:
		m.coordinate = m.coordinate.Add(&util.Coordinate{0, -1})
	case EAST:
		m.coordinate = m.coordinate.Add(&util.Coordinate{1, 0})
	case WEST:
		m.coordinate = m.coordinate.Add(&util.Coordinate{-1, 0})
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input value>\n", os.Args[0])
		os.Exit(1)
	}

	value, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}

	dirCount := 0
	distance := 0
	directions := []Direction{EAST, NORTH, WEST, SOUTH}
	mover := &Mover{&util.Coordinate{0, 0}, 0}

	for i := 0; ; i++ {
		if i%2 == 0 {
			distance++
		}
		direction := directions[i%4]
		for j := 0; j < distance; j++ {
			mem := 0
			mover.coordinate.Neighbors(func(c *util.Coordinate) {
				if v, found := values[c.String()]; found {
					mem += v
				}
			})
			mover.coordinate.Diagonals(func(c *util.Coordinate) {
				if v, found := values[c.String()]; found {
					mem += v
				}
			})

			if mem == 0 {
				mem = 1
			}
			values[mover.coordinate.String()] = mem

			if mem > value {
				fmt.Printf("Mem value is %d\n", mem)
				os.Exit(0)
			}

			if mover.id == value {
				fmt.Printf("Now at %s\n", mover.coordinate)
				fmt.Printf("Distance to get to %d is %d\n", value, util.ManhattanDistance(mover.coordinate, &util.Coordinate{0, 0}))
				os.Exit(0)
			}

			mover.Move(direction)
			dirCount++
		}
	}
}
