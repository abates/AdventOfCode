package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
)

func main() {
	sum := 0
	for _, line := range util.ReadInput() {
		room := parseRoom(line)
		if room.Valid() {
			sum += room.sectorId
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}
