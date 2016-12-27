package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/alg"
	"github.com/abates/AdventOfCode/2016/util"
	"strings"
)

func part1(grid *Grid) {
	count := 0
	grid.Iterate(func(l *Node) {
		grid.Iterate(func(r *Node) {
			if l != r && l.used > 0 && l.used <= r.avail {
				count++
			}
		})
	})

	fmt.Printf("Number of pairs: %d\n", count)
}

func part2(grid *Grid) {
	var endState *Grid
	var p []alg.Node
	alg.Traverse(grid, func(level int, path []alg.Node) bool {
		p = path
		position := path[len(path)-1]
		if grid, ok := position.(*Grid); ok {
			endState = grid
			return grid.free.X == 37 && grid.free.Y == 0
		}
		return false
	})
	fmt.Printf("%v %v\n", len(p)-1, endState)

	grid = NewGrid()
	grid.grid[0] = endState.grid[0]
	grid.grid[1] = endState.grid[1]
	grid.free = endState.free
	grid.data = endState.data

	alg.Traverse(grid, func(level int, path []alg.Node) bool {
		p = path
		position := path[len(path)-1]
		if grid, ok := position.(*Grid); ok {
			endState = grid
			return grid.data.X == 0 && grid.data.Y == 0
		}
		return false
	})

	fmt.Printf("%v %v\n", len(p)-1, endState)
}

func main() {
	grid := NewGrid()
	for _, line := range util.ReadInput() {
		if strings.HasPrefix(line, "/dev") {
			grid.AddNode(NewNodeFromString(line))
		}
	}
	part1(grid)
	part2(grid)

}
