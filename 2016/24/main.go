package main

import (
	"sort"
	"strings"

	"github.com/abates/AdventOfCode/2016/alg"
	"github.com/abates/AdventOfCode/2016/util"
	"github.com/abates/AdventOfCode/graph"
	"github.com/cznic/mathutil"
)

type TerminalExplorer struct {
	*alg.Coordinate
	name string
}

func (e *TerminalExplorer) Name() string {
	return e.name
}

func (e *TerminalExplorer) Neighbors() []alg.Node {
	return nil
}

type Explorer struct {
	position *alg.Coordinate
	matrix   *Matrix
	name     string
}

func (e *Explorer) Name() string {
	return e.name
}

func (e *Explorer) ID() string {
	return e.position.ID()
}

func (e *Explorer) Neighbors() []alg.Node {
	neighbors := make([]alg.Node, 0)
	maxX := len(e.matrix.walls[0]) - 1
	maxY := len(e.matrix.walls) - 1
	for _, node := range e.position.Neighbors() {
		if coordinate, ok := node.(*alg.Coordinate); ok {
			if coordinate.Get(0) < 0 || coordinate.Get(1) < 0 || coordinate.Get(0) > maxX || coordinate.Get(1) > maxY {
				continue
			}

			if e.matrix.IsNode(coordinate) {
				neighbors = append(neighbors, &TerminalExplorer{coordinate, e.matrix.GetNode(coordinate).Name()})
			} else if !e.matrix.walls[coordinate.Get(1)][coordinate.Get(0)] {
				neighbors = append(neighbors, &Explorer{coordinate, e.matrix, ""})
			}
		}
	}
	return neighbors
}

type Matrix struct {
	walls map[int]map[int]bool
	nodes map[int]map[int]*Explorer
}

func (m *Matrix) IsNode(coordinate *alg.Coordinate) bool {
	if _, found := m.nodes[coordinate.Get(1)]; found {
		_, found = m.nodes[coordinate.Get(1)][coordinate.Get(0)]
		return found
	}
	return false
}

func (m *Matrix) GetNode(coordinate *alg.Coordinate) *Explorer {
	if _, found := m.nodes[coordinate.Get(1)]; found {
		return m.nodes[coordinate.Get(1)][coordinate.Get(0)]
	}
	return nil
}

func main() {
	matrix := &Matrix{
		walls: make(map[int]map[int]bool),
		nodes: make(map[int]map[int]*Explorer),
	}

	for y, line := range util.ReadInput() {
		matrix.walls[y] = make(map[int]bool)
		for x, t := range strings.Split(line, "") {
			switch t {
			case "#":
				matrix.walls[y][x] = true
			case ".":
				matrix.walls[y][x] = false
			default:
				matrix.walls[y][x] = false
				if _, found := matrix.nodes[y]; !found {
					matrix.nodes[y] = make(map[int]*Explorer)
				}
				matrix.nodes[y][x] = &Explorer{alg.NewCoordinate(x, y), matrix, t}
			}
		}
	}

	g := &graph.BasicGraph{}
	nodeIds := make([]string, 0)
	for _, row := range matrix.nodes {
		for _, explorer := range row {
			nodeIds = append(nodeIds, explorer.Name())
			alg.Traverse(explorer, func(l int, path []alg.Node) bool {
				switch neighbor := path[len(path)-1].(type) {
				case *TerminalExplorer:
					g.AddDirectedEdge(explorer.Name(), matrix.GetNode(neighbor.Coordinate).Name(), len(path)-1)
				case *Explorer:
				}
				return false
			})
		}
	}

	distances := graph.SPFAll(g)

	min := 0
	ids := sort.StringSlice(nodeIds)
	for mathutil.PermutationFirst(ids); mathutil.PermutationNext(ids); {
		node := "0"
		distance := 0
		for _, id := range ids {
			distance += distances[g.GetNode(node)][g.GetNode(id)]
			node = id
		}
		distance += distances[g.GetNode(node)][g.GetNode("0")]

		if min == 0 || distance < min {
			min = distance
		}
	}
	println(min)
}
