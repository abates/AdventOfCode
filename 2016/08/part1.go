package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
	"strings"
)

func main() {
	screen := NewScreen(50, 6)
	for _, line := range util.ReadInput() {
		tokens := strings.SplitN(line, " ", 2)
		if tokens[0] == "rect" {
			rows := 0
			columns := 0
			fmt.Sscanf(tokens[1], "%dx%d", &rows, &columns)
			screen.Rect(rows, columns)
		} else if tokens[0] == "rotate" {
			tokens = strings.SplitN(tokens[1], " ", 2)
			location := 0
			amount := 0
			if tokens[0] == "row" {
				fmt.Sscanf(tokens[1], "y=%d by %d", &location, &amount)
				screen.RotateRow(location, amount)
			} else if tokens[0] == "column" {
				fmt.Sscanf(tokens[1], "x=%d by %d", &location, &amount)
				screen.RotateColumn(location, amount)
			} else {
				fmt.Printf("Don't know how to rotate %s\n", tokens[0])
			}
		} else {
			fmt.Printf("Don't know how to handle %s\n", tokens[0])
		}
	}

	count := 0
	for _, row := range screen.pixels {
		for _, column := range row {
			if column == "#" {
				count++
			}
		}
	}
	fmt.Printf("%d\n%s\n", count, screen)
}
