package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
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
			walker.Walk(distance)
			return true
		})
		if err != nil {
			break
		}
		fmt.Printf("Distance: %d\n", walker.position.Distance(Coordinate{0, 0}))
	}
}
