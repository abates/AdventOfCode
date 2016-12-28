package main

import (
	"github.com/abates/AdventOfCode/2016/alg"
	"github.com/abates/AdventOfCode/2016/util"
	"github.com/cznic/mathutil"
	"sort"
	"strings"
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
			if coordinate.X < 0 || coordinate.Y < 0 || coordinate.X > maxX || coordinate.Y > maxY {
				continue
			}

			if e.matrix.IsNode(coordinate) {
				neighbors = append(neighbors, &TerminalExplorer{coordinate, e.matrix.GetNode(coordinate).Name()})
			} else if !e.matrix.walls[coordinate.Y][coordinate.X] {
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
	if _, found := m.nodes[coordinate.Y]; found {
		_, found = m.nodes[coordinate.Y][coordinate.X]
		return found
	}
	return false
}

func (m *Matrix) GetNode(coordinate *alg.Coordinate) *Explorer {
	if _, found := m.nodes[coordinate.Y]; found {
		return m.nodes[coordinate.Y][coordinate.X]
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
				matrix.nodes[y][x] = &Explorer{&alg.Coordinate{x, y}, matrix, t}
			}
		}
	}

	graph := alg.NewBasicGraph()
	nodeIds := make([]string, 0)
	for _, row := range matrix.nodes {
		for _, explorer := range row {
			nodeIds = append(nodeIds, explorer.Name())
			alg.Traverse(explorer, func(l int, path []alg.Node) bool {
				switch neighbor := path[len(path)-1].(type) {
				case *TerminalExplorer:
					node := graph.GetNode(explorer.Name())
					if node == nil {
						node = alg.NewBasicGraphNode(explorer.Name())
						graph.AddNode(node)
					}
					node.(*alg.BasicGraphNode).AddEdge(alg.NewBasicEdge(len(path)-1, matrix.GetNode(neighbor.Coordinate).Name()))
				case *Explorer:
				}
				return false
			})
		}
	}

	distances := alg.SPFAll(graph)

	min := 0
	ids := sort.StringSlice(nodeIds)
	for mathutil.PermutationFirst(ids); mathutil.PermutationNext(ids); {
		node := "0"
		distance := 0
		for _, id := range ids {
			distance += distances[node][id]
			node = id
		}
		distance += distances[node]["0"]

		if min == 0 || distance < min {
			min = distance
		}
	}
	println(min)
}
