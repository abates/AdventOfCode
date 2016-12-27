package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/alg"
)

func main() {
	distance := MagicWalk(&alg.Coordinate{1, 1}, &alg.Coordinate{31, 39}, 1364)
	fmt.Printf("Distance %d\n", distance-1)

	visited := MagicWalkMax(&alg.Coordinate{1, 1}, 1364, 50)
	fmt.Printf("Num Visited: %d\n", visited)
}
