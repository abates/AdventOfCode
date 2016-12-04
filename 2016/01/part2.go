package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
	"os"
)

func main() {
	for _, line := range util.ReadInput() {
		visited := make(map[string]bool)
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
			walker.Walk(operation.distance, func(c Coordinate) {
				if _, found := visited[walker.position.String()]; found {
					fmt.Printf("Distance %d\n", walker.position.Distance(Coordinate{0, 0}))
					os.Exit(0)
				} else {
					visited[walker.position.String()] = true
				}
			})
		}
		fmt.Printf("Distance: %d\n", walker.position.Distance(Coordinate{0, 0}))
	}
}
