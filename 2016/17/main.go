package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/bfs"
)

func main() {
	id := ""
	bfs.Traverse(NewVaultCoordinate(1, 1, "qtetzkpl"), func(l int, p []bfs.Node) bool {
		c := p[len(p)-1].(*VaultCoordinate)
		id = c.ID()
		return c.Coordinate.ID() == "(4,4)"
	})
	fmt.Printf("ID: %v\n", id)
	fmt.Printf("Height: %d\n", bfs.Height(NewVaultCoordinate(1, 1, "qtetzkpl"))-1)
}
