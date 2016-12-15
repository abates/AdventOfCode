package main

import (
	"fmt"
)

func main() {
	walker := NewMagicWalker(1, 1, 1364)

	distance := walker.Walk(Coordinate{31, 39})
	//fmt.Printf("%s\n", Draw(50, 50, MagicDetector(1364)))
	fmt.Printf("Distance %d\n", distance)

	walker = NewMagicWalker(1, 1, 1364)
	numVisited := walker.WalkMax(Coordinate{31, 39}, 50)
	fmt.Printf("Num Visited: %d\n", numVisited)
}
