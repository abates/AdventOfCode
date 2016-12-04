package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Operation struct {
	direction string
	distance  int
}

func splitLine(line string) []Operation {
	operations := make([]Operation, 0)

	for _, s := range strings.Split(line, ", ") {
		tokens := strings.SplitN(s, "", 2)
		distance, _ := strconv.Atoi(tokens[1])
		operations = append(operations, Operation{tokens[0], distance})
	}

	return operations
}

type Coordinate struct {
	x int
	y int
}

func (c Coordinate) Distance(other Coordinate) int {
	return int(math.Abs(float64(other.x-c.x)) + math.Abs(float64(other.y-c.y)))
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d,%d)", c.x, c.y)
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Walker struct {
	position  Coordinate
	direction Direction
}

func (w *Walker) TurnLeft() {
	w.direction -= 1
	if w.direction < North {
		w.direction = West
	}
}

func (w *Walker) TurnRight() {
	w.direction += 1
	if w.direction > West {
		w.direction = North
	}
}

func (w *Walker) Walk(distance int, cbs ...func(Coordinate)) {
	for i := 0; i < distance; i++ {
		switch w.direction {
		case North:
			w.position.y += 1
		case South:
			w.position.y -= 1
		case East:
			w.position.x += 1
		case West:
			w.position.x -= 1
		}

		if len(cbs) > 0 {
			for _, cb := range cbs {
				cb(w.position)
			}
		}
	}
}
