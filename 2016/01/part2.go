package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		visited := make(map[string]bool)
		walker := &Walker{}
		err := readInput(reader, func(operation string, distance int) bool {
			switch operation {
			case "L":
				walker.TurnLeft()
			case "R":
				walker.TurnRight()
			default:
				println("Unknown operation", operation)
			}
			done := false
			walker.Walk(distance, func(c Coordinate) {
				if _, found := visited[walker.position.String()]; found {
					fmt.Printf("Distance %d\n", walker.position.Distance(Coordinate{0, 0}))
					done = true
				} else {
					visited[walker.position.String()] = true
				}
			})
			return !done
		})
		if err != nil {
			break
		}
	}
}
