package main

import (
	"fmt"
)

func main() {
	count := 0
	for _, row := range readInput() {
		if isTriangle(row[0], row[1], row[2]) {
			count += 1
		}
	}
	fmt.Printf("Count: %d\n", count)
}
