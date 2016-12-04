package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
)

func main() {
	for _, line := range util.ReadInput() {
		walker := &Walker{}
		for _, operation := range splitLine(line) {
			switch operation.direction {
			case "L":
				walker.TurnLeft()
			case "R":
				walker.TurnRight()
			default:
				println("Unknown direction", operation.direction)
			}
			walker.Walk(operation.distance)
		}
		fmt.Printf("Distance: %d\n", walker.position.Distance(Coordinate{0, 0}))
	}
}
