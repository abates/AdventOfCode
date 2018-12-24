package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/abates/AdventOfCode/graph"
	"github.com/abates/AdventOfCode/graph/bfs"
)

type Tool int

const (
	ToolNeither Tool = iota
	ToolTorch
	ToolClimbing
)

type Edge struct {
	neighbor *Node
}

func (e *Edge) Weight() int          { return 1 }
func (e *Edge) Neighbor() graph.Node { return e.neighbor }

type Node struct {
	x     int
	y     int
	wait  int
	tool  Tool
	scan  *Scan
	edges []graph.Edge
}

func (n *Node) Edges() []graph.Edge {
	if len(n.edges) > 0 {
		return n.edges
	}

	if n.wait > 0 {
		n.edges = []graph.Edge{&Edge{n.scan.lookupNode(n.x, n.y, n.tool, n.wait-1)}}
		return n.edges
	}

	for _, delta := range [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		deltaX := delta[0]
		deltaY := delta[1]
		if n.x+deltaX < 0 || n.y+deltaY < 0 {
			continue
		}

		neighborSoil := n.scan.SoilType(n.x+deltaX, n.y+deltaY)
		neighborTools := []Tool{}
		// 0 for rocky regions, 1 for wet regions, and 2 for narrow regions
		switch neighborSoil {
		case 0: // rocky
			neighborTools = []Tool{ToolClimbing, ToolTorch}
		case 1: // wet
			neighborTools = []Tool{ToolNeither, ToolClimbing}
		case 2: // narrow
			neighborTools = []Tool{ToolNeither, ToolTorch}
		}

		for _, tool := range neighborTools {
			if tool == n.tool {
				neighbor := n.scan.lookupNode(n.x+deltaX, n.y+deltaY, tool, 0)
				n.edges = append(n.edges, &Edge{neighbor})
			} else {
				neighbor := n.scan.lookupNode(n.x+deltaX, n.y+deltaY, tool, 7)
				n.edges = append(n.edges, &Edge{neighbor})
			}
		}
	}
	return n.edges
}

type Graph struct {
	Start *Node
	End   *Node
}

type Scan struct {
	eroLevel map[[2]int]int
	nodes    map[[4]int]*Node
	target   [2]int
	depth    int
}

func (s *Scan) lookupNode(x, y int, tool Tool, wait int) *Node {
	node, found := s.nodes[[4]int{x, y, int(tool), wait}]
	if !found {
		if s.nodes == nil {
			s.nodes = make(map[[4]int]*Node)
		}
		node = &Node{x: x, y: y, tool: tool, wait: wait, scan: s}
		s.nodes[[4]int{x, y, int(tool), wait}] = node
	}
	return node
}

func (s *Scan) UnmarshalText(text []byte) (err error) {
	for _, line := range bytes.Split(text, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		str := string(line)
		if i := strings.Index(str, "depth: "); i >= 0 {
			_, err = fmt.Sscanf(str, "depth: %d", &s.depth)
		} else if i := strings.Index(str, "target: "); i >= 0 {
			_, err = fmt.Sscanf(str, "target: %d,%d", &s.target[0], &s.target[1])
		}

		if err != nil {
			break
		}
	}

	return
}

func (s *Scan) ErosionLevel(x, y int) int {
	el, found := s.eroLevel[[2]int{x, y}]
	if !found {
		if s.eroLevel == nil {
			s.eroLevel = make(map[[2]int]int)
		}
		// geologic index
		gi := 0
		if (x == 0 && y == 0) || (x == s.target[0] && y == s.target[1]) {
			gi = 0
		} else if y == 0 {
			gi = x * 16807
		} else if x == 0 {
			gi = y * 48271
		} else {
			gi = s.ErosionLevel(x-1, y) * s.ErosionLevel(x, y-1)
		}
		// erosion level
		el = (gi + s.depth) % 20183
		s.eroLevel[[2]int{x, y}] = el
	}
	return el
}

func (s *Scan) SoilType(x, y int) int {
	return s.ErosionLevel(x, y) % 3
}

func (s *Scan) RiskLevel() int {
	riskLevel := 0
	for x := 0; x <= s.target[0]; x++ {
		for y := 0; y <= s.target[1]; y++ {
			riskLevel += s.SoilType(x, y)
		}
	}
	return riskLevel
}

func (s *Scan) String() string {
	var builder strings.Builder
	for x := 0; x <= s.target[0]+1; x++ {
		for y := 0; y <= s.target[1]+1; y++ {
			if x == 0 && y == 0 {
				builder.WriteString("M")
			} else if x == s.target[0] && y == s.target[1] {
				builder.WriteString("T")
			} else {
				soilType := s.SoilType(x, y)
				if soilType == 0 {
					builder.WriteString(".")
				} else if soilType == 1 {
					builder.WriteString("=")
				} else if soilType == 2 {
					builder.WriteString("|")
				} else {
					panic(fmt.Sprintf("Unknown soil type: %d", soilType))
				}
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (s *Scan) BuildGraph() *Graph {
	graph := &Graph{}
	return graph
}

func part2(input []byte) error {
	scan := &Scan{}
	err := scan.UnmarshalText(input)
	if err == nil {
		start := scan.lookupNode(0, 0, ToolTorch, 0)
		end := scan.lookupNode(scan.target[0], scan.target[1], ToolTorch, 0)
		path := bfs.Find(start, end)
		fmt.Printf("Part 2: %d\n", len(path)-1)
	}
	return err
}

func part1(input []byte) error {
	scan := &Scan{}
	err := scan.UnmarshalText(input)
	if err == nil {
		riskLevel := scan.RiskLevel()
		fmt.Printf("Part 1: %d\n", riskLevel)
	}
	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <input file>\n", os.Args[0])
		os.Exit(-1)
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input file: %v\n", err)
		os.Exit(-1)
	}

	for i, f := range []func([]byte) error{part1, part2} {
		err = f(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Part %d failed: %v\n", i, err)
			os.Exit(-1)
		}
	}
}
