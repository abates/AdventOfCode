package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/abates/AdventOfCode/graph"
)

type Edge struct {
	neighbor graph.Node
}

func (e *Edge) Weight() int          { return 1 }
func (e *Edge) Neighbor() graph.Node { return e.neighbor }

type Node struct {
	edges map[[2]int]*Edge
}

func (n *Node) Edges() []graph.Edge {
	edges := []graph.Edge{}
	for _, edge := range n.edges {
		edges = append(edges, edge)
	}
	return edges
}

type Map struct {
	start *Node
	index map[[2]int]*Node
}

func (m *Map) Nodes() (nodes []graph.Node) {
	for _, node := range m.index {
		nodes = append(nodes, node)
	}
	return nodes
}

func (m *Map) AddEdge(from, to [2]int) {
	m.addEdge(to, from)
	m.addEdge(from, to)
}

func (m *Map) addEdge(from, to [2]int) {
	if m.index == nil {
		m.index = make(map[[2]int]*Node)
	}

	fromNode, found := m.index[from]
	if !found {
		fromNode = &Node{make(map[[2]int]*Edge)}
		m.index[from] = fromNode
	}

	toNode, found := m.index[to]
	if !found {
		toNode = &Node{make(map[[2]int]*Edge)}
		m.index[to] = toNode
	}

	if _, found := fromNode.edges[to]; !found {
		fromNode.edges[to] = &Edge{neighbor: toNode}
	}
}

func (m *Map) UnmarshalText(input []byte) (err error) {
	return m.parse(strings.Split(string(input), ""))
}

func (m *Map) parse(str []string) error {
	directions := map[string][2]int{
		"N": {0, -1},
		"S": {0, 1},
		"E": {1, 0},
		"W": {-1, 0},
	}
	starts := [][2]int{{0, 0}}

	position := starts[0]
	for _, s := range str {
		if strings.Contains("NSEW", s) {
			dir := directions[s]
			newPosition := [2]int{position[0] + dir[0], position[1] + dir[1]}
			// add edge from previous position to new position
			m.AddEdge(position, newPosition)
			position = newPosition
		} else if s == "$" {
			break
		} else if s == "(" {
			starts = append([][2]int{position}, starts...)
		} else if s == ")" {
			starts = starts[1:]
		} else if s == "|" {
			// record the end of this pattern
			starts = append(starts, position)
			// rewind to where we were at the beginning of the group
			position = starts[0]
		}
	}
	return nil
}

func part1(input []byte) error {
	m := &Map{}
	err := m.UnmarshalText(input)
	if err == nil {
		v := graph.SPF(m, m.index[[2]int{0, 0}])
		max := 0
		for _, d := range v {
			if max < d {
				max = d
			}
		}
		fmt.Printf("Part 1: %d\n", max)
	}
	return err
}

func part2(input []byte) error {
	m := &Map{}
	err := m.UnmarshalText(input)
	if err == nil {
		v := graph.SPF(m, m.index[[2]int{0, 0}])
		sum := 0
		for _, d := range v {
			if d >= 1000 {
				sum++
			}
		}
		fmt.Printf("Part 2: %d\n", sum)
	}
	return err
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(fmt.Sprintf("could not create CPU profile: ", err))
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(fmt.Sprintf("could not start CPU profile: ", err))
	}

	defer pprof.StopCPUProfile()

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input file>\n", os.Args[0])
		os.Exit(-1)
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input file: %v\n", err)
		os.Exit(-1)
	}

	for i, f := range []func([]byte) error{part1, part2} {
		err := f(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Part %d failed: %v\n", i+1, err)
			os.Exit(-1)
		}
	}
}
