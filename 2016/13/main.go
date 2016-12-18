package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
)

func main() {
	walker := NewMagicWalker(1, 1, 1364)

	distance := walker.Walk(&util.Coordinate{31, 39})
	fmt.Printf("Distance %d\n", distance-1)

	walker = NewMagicWalker(1, 1, 1364)
	numVisited := walker.WalkMax(50)
	fmt.Printf("Num Visited: %d\n", numVisited)
}
