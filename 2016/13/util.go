package main

import (
	"github.com/abates/AdventOfCode/2016/util"
)

type WallDetector func(*util.Coordinate) bool

func MagicDetector(magicNumber int) WallDetector {
	return func(coordinate *util.Coordinate) bool {
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
	*util.Coordinate
	isOpenSpace WallDetector
}

func (m *MazeCoordinate) Equal(node util.Node) bool {
	if other, ok := node.(*MazeCoordinate); ok {
		return m.Coordinate.Equal(other.Coordinate)
	}
	return false
}

func (m *MazeCoordinate) Neighbors() []util.Node {
	nodes := make([]util.Node, 0)
	for _, candidate := range m.Coordinate.Neighbors() {
		if coordinate, ok := candidate.(*util.Coordinate); ok {
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
			if detector(&util.Coordinate{x, y}) {
				writer.Write(".")
			} else {
				writer.Write("#")
			}
		}
		writer.Write("\n")
	}

	return writer.String()
}

type Walker struct {
	tree *util.Tree
}

func NewWalker(startX, startY int, wallDetector WallDetector) *Walker {
	return &Walker{
		tree: util.NewTree(&MazeCoordinate{&util.Coordinate{startX, startY}, wallDetector}),
	}
}

func NewMagicWalker(startX, startY, magicNumber int) *Walker {
	return NewWalker(startX, startY, MagicDetector(magicNumber))
}

func (w *Walker) Walk(destination *util.Coordinate) int {
	return len(w.tree.Find(&MazeCoordinate{destination, nil}))
}

func (w *Walker) WalkMax(maxSteps int) int {
	w.tree.FindAt(maxSteps)
	visited := w.tree.VisitedNodes()
	return len(visited)
}
