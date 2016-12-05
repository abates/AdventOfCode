package main

import (
	"fmt"
)

func main() {
	count := 0
	rows := readInput()
	for i := 0; i < len(rows); i += 3 {
		for j := 0; j < 3; j++ {
			if isTriangle(rows[i][j], rows[i+1][j], rows[i+2][j]) {
				count += 1
			}
		}
	}
	fmt.Printf("Count: %d\n", count)
}
