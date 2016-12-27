package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/bfs"
)

func main() {
	distance := MagicWalk(&bfs.Coordinate{1, 1}, &bfs.Coordinate{31, 39}, 1364)
	fmt.Printf("Distance %d\n", distance-1)

	visited := MagicWalkMax(&bfs.Coordinate{1, 1}, 1364, 50)
	fmt.Printf("Num Visited: %d\n", visited)
}
