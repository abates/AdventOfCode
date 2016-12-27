package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/alg"
)

func main() {
	id := ""
	alg.Traverse(NewVaultCoordinate(1, 1, "qtetzkpl"), func(l int, p []alg.Node) bool {
		c := p[len(p)-1].(*VaultCoordinate)
		id = c.ID()
		return c.Coordinate.ID() == "(4,4)"
	})
	fmt.Printf("ID: %v\n", id)
	fmt.Printf("Height: %d\n", alg.Height(NewVaultCoordinate(1, 1, "qtetzkpl"))-1)
}
