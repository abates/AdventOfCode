package main

import (
	"crypto/md5"
	"fmt"
	"github.com/abates/AdventOfCode/2016/bfs"
)

type VaultCoordinate struct {
	*bfs.Coordinate
	path string
}

func NewVaultCoordinate(x, y int, path string) *VaultCoordinate {
	return &VaultCoordinate{
		Coordinate: &bfs.Coordinate{x, y},
		path:       path,
	}
}

func (v *VaultCoordinate) ID() string {
	return v.path
}

func (v *VaultCoordinate) Equal(node bfs.Node) bool {
	if other, ok := node.(*VaultCoordinate); ok {
		return v.Coordinate.Equal(other.Coordinate)
	}
	return false
}

func (v *VaultCoordinate) Neighbors() []bfs.Node {
	if v.X == 4 && v.Y == 4 {
		return nil
	}

	neighbors := make([]bfs.Node, 0)
	directions := []*bfs.Coordinate{
		&bfs.Coordinate{0, -1},
		&bfs.Coordinate{0, 1},
		&bfs.Coordinate{-1, 0},
		&bfs.Coordinate{1, 0},
	}

	sum := md5.Sum([]byte(v.path))
	for direction, delta := range directions {
		coordinate := v.Coordinate.Add(delta)
		if coordinate.X <= 0 || coordinate.X > 4 || coordinate.Y <= 0 || coordinate.Y > 4 {
			continue
		}
		open := byte(0x00)
		dstr := ""
		switch direction {
		case 0:
			dstr = "U"
			open = sum[0] >> 4
		case 1:
			dstr = "D"
			open = sum[0] & 0x0f
		case 2:
			dstr = "L"
			open = sum[1] >> 4
		case 3:
			dstr = "R"
			open = sum[1] & 0x0f
		}

		if open > 0x0a {
			neighbors = append(neighbors, NewVaultCoordinate(coordinate.X, coordinate.Y, fmt.Sprintf("%s%s", v.path, dstr)))
		}
	}
	return neighbors
}

func (v *VaultCoordinate) String() string {
	return v.ID()
}
