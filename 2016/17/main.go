package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
)

func main() {
	tree := util.NewTree(NewVaultCoordinate(1, 1, "qtetzkpl"))
	//tree := util.NewTree(NewVaultCoordinate(1, 1, "ihgpwlah"))
	path := tree.Find(NewVaultCoordinate(4, 4, ""))
	fmt.Printf("ID: %s\n", path[len(path)-1].ID())

	tree = util.NewTree(NewVaultCoordinate(1, 1, "qtetzkpl"))
	fmt.Printf("Height: %d\n", tree.Height(NewVaultCoordinate(4, 4, ""))-1)
}
