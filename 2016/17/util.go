package main

import (
	"crypto/md5"
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
)

type VaultCoordinate struct {
	*util.Coordinate
	path string
}

func NewVaultCoordinate(x, y int, path string) *VaultCoordinate {
	return &VaultCoordinate{
		Coordinate: &util.Coordinate{x, y},
		path:       path,
	}
}

func (v *VaultCoordinate) ID() string {
	return v.path
}

func (v *VaultCoordinate) Equal(node util.Node) bool {
	if other, ok := node.(*VaultCoordinate); ok {
		return v.Coordinate.Equal(other.Coordinate)
	}
	return false
}

func (v *VaultCoordinate) Neighbors() []util.Node {
	if v.X == 4 && v.Y == 4 {
		return nil
	}

	neighbors := make([]util.Node, 0)
	directions := map[string]*util.Coordinate{
		"U": &util.Coordinate{0, -1},
		"D": &util.Coordinate{0, 1},
		"L": &util.Coordinate{-1, 0},
		"R": &util.Coordinate{1, 0},
	}

	sum := md5.Sum([]byte(v.path))
	for direction, delta := range directions {
		coordinate := v.Coordinate.Add(delta)
		if coordinate.X <= 0 || coordinate.X > 4 || coordinate.Y <= 0 || coordinate.Y > 4 {
			continue
		}
		open := byte(0x00)
		switch direction {
		case "U":
			open = sum[0] >> 4
		case "D":
			open = sum[0] & 0x0f
		case "L":
			open = sum[1] >> 4
		case "R":
			open = sum[1] & 0x0f
		}

		if open > 0x0a {
			neighbors = append(neighbors, NewVaultCoordinate(coordinate.X, coordinate.Y, fmt.Sprintf("%s%s", v.path, direction)))
		}
	}
	return neighbors
}

func (v *VaultCoordinate) String() string {
	return v.ID()
}
