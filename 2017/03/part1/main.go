// I hate this solution... :-/
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abates/AdventOfCode/coordinate"
)

var values = make(map[string]int)

type Direction *coordinate.Coordinate

var (
	North     = Direction(coordinate.New(0, 1))
	Northeast = Direction(coordinate.New(1, 1))
	East      = Direction(coordinate.New(1, 0))
	Southeast = Direction(coordinate.New(1, -1))
	South     = Direction(coordinate.New(0, -1))
	Southwest = Direction(coordinate.New(-1, -1))
	West      = Direction(coordinate.New(-1, 0))
	Northwest = Direction(coordinate.New(-1, 1))
)

type Mover struct {
	coordinate *coordinate.Coordinate
	id         int
}

func (m *Mover) Move(direction Direction) {
	m.coordinate = m.coordinate.Add(direction)
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
	directions := []Direction{East, North, West, South}
	mover := &Mover{coordinate.New(0, 0), 0}

	for i := 0; ; i++ {
		if i%2 == 0 {
			distance++
		}
		direction := directions[i%4]
		for j := 0; j < distance; j++ {
			mem := 0
			for _, direction := range []Direction{East, Northeast, North, Northwest, West, Southwest, South, Southeast} {
				c := mover.coordinate.Add(direction)
				if v, found := values[c.String()]; found {
					mem += v
				}
			}

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
				fmt.Printf("Distance to get to %d is %d\n", value, coordinate.ManhattanDistance(mover.coordinate, coordinate.New(0, 0)))
				os.Exit(0)
			}

			mover.Move(direction)
			dirCount++
		}
	}
}
