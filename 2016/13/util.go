package main

import (
	"github.com/abates/AdventOfCode/2016/alg"
	"github.com/abates/AdventOfCode/2016/util"
)

type WallDetector func(*alg.Coordinate) bool

func MagicDetector(magicNumber int) WallDetector {
	return func(coordinate *alg.Coordinate) bool {
		value := coordinate.X*coordinate.X + 3*coordinate.X + 2*coordinate.X*coordinate.Y + coordinate.Y + coordinate.Y*coordinate.Y + magicNumber
		count := 0
		for value > 0 {
			if value&0x01 == 1 {
				count++
			}
			value = value >> 1
		}

		return count%2 == 0
	}
}

type MazeCoordinate struct {
	*alg.Coordinate
	isOpenSpace WallDetector
}

func (m *MazeCoordinate) Equal(node alg.Node) bool {
	if other, ok := node.(*MazeCoordinate); ok {
		return m.Coordinate.Equal(other.Coordinate)
	}
	return false
}

func (m *MazeCoordinate) Neighbors() []alg.Node {
	nodes := make([]alg.Node, 0)
	for _, candidate := range m.Coordinate.Neighbors() {
		if coordinate, ok := candidate.(*alg.Coordinate); ok {
			if coordinate.X < 0 || coordinate.Y < 0 {
				continue
			}
			if m.isOpenSpace(coordinate) {
				nodes = append(nodes, &MazeCoordinate{coordinate, m.isOpenSpace})
			}
		}
	}
	return nodes
}

func Draw(width, height int, detector WallDetector) string {
	writer := util.StringWriter{}
	writer.Write("  ")
	for x := 0; x < width; x++ {
		writer.Writef("%d", x)
	}
	writer.Write("\n")

	for y := 0; y < height; y++ {
		writer.Writef("%d ", y)
		for x := 0; x < width; x++ {
			if detector(&alg.Coordinate{x, y}) {
				writer.Write(".")
			} else {
				writer.Write("#")
			}
		}
		writer.Write("\n")
	}

	return writer.String()
}

func MagicWalk(origin, destination *alg.Coordinate, magicNumber int) int {
	return Walk(origin, destination, MagicDetector(magicNumber))
}

func Walk(origin, destination *alg.Coordinate, wallDetector WallDetector) int {
	rootNode := &MazeCoordinate{
		Coordinate:  origin,
		isOpenSpace: wallDetector,
	}
	return len(alg.Find(rootNode, destination.ID()))
}

func MagicWalkMax(origin *alg.Coordinate, magicNumber, maxSteps int) int {
	return WalkMax(origin, MagicDetector(magicNumber), maxSteps)
}

func WalkMax(origin *alg.Coordinate, wallDetector WallDetector, maxSteps int) int {
	rootNode := &MazeCoordinate{
		Coordinate:  origin,
		isOpenSpace: wallDetector,
	}

	visited := make(map[string]bool)
	alg.Traverse(rootNode, func(l int, p []alg.Node) bool {
		if p[len(p)-1].ID() != rootNode.ID() {
			visited[p[len(p)-1].ID()] = true
		}

		if l == maxSteps+1 {
			return true
		}
		return false
	})

	return len(visited)
}
