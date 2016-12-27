package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/bfs"
	"strconv"
	"strings"
)

func convertToSize(str string) (size int) {
	if strings.HasSuffix(str, "T") {
		size, _ = strconv.Atoi(str[:len(str)-1])
	}
	return size
}

func convertToPercent(str string) (percent float32) {
	p := 0
	if strings.HasSuffix(str, "%") {
		p, _ = strconv.Atoi(str[:len(str)-1])
		percent = float32(p) / 100
	}
	return percent
}

type Node struct {
	coordinate *bfs.Coordinate
	filesystem string
	size       int
	used       int
	avail      int
	use        float32
}

func NewNodeFromString(str string) *Node {
	fields := strings.Fields(str)
	var x, y int
	fmt.Sscanf(fields[0], "/dev/grid/node-x%d-y%d", &x, &y)
	node := &Node{
		coordinate: &bfs.Coordinate{x, y},
		filesystem: fields[0],
		size:       convertToSize(fields[1]),
		used:       convertToSize(fields[2]),
		avail:      convertToSize(fields[3]),
		use:        convertToPercent(fields[4]),
	}
	return node
}

func (n *Node) Copy() *Node {
	newNode := &Node{
		filesystem: n.filesystem,
		size:       n.size,
		used:       n.used,
		avail:      n.avail,
		use:        n.use,
	}

	newNode.coordinate = &bfs.Coordinate{n.coordinate.X, n.coordinate.Y}
	return newNode
}

func (n *Node) MoveTo(destination *Node) {
	destination.used += n.used
	n.used = 0
	destination.avail = destination.size - destination.used
	n.avail = n.size
	destination.use = float32(destination.used) / float32(destination.size)
	n.use = 0
}

func (n *Node) Neighbors() []*bfs.Coordinate {
	neighbors := make([]*bfs.Coordinate, 0)
	for _, c := range n.coordinate.Neighbors() {
		coordinate := c.(*bfs.Coordinate)
		if coordinate.X < 0 || coordinate.Y < 0 {
			continue
		}
		neighbors = append(neighbors, c.(*bfs.Coordinate))
	}
	return neighbors
}

func (n *Node) ID() string {
	return n.coordinate.ID()
}

func (n *Node) Equal(other *Node) bool {
	return n.coordinate.Equal(other.coordinate)
}

func (n *Node) String() string {
	filesystem := fmt.Sprintf("/dev/grid/node-x%d-y%d", n.coordinate.X, n.coordinate.Y)
	return fmt.Sprintf("%-23s %3dT %4dT  %4dT   %.0f%%", filesystem, n.size, n.used, n.avail, n.use*100)
}

type Grid struct {
	data *bfs.Coordinate
	free *bfs.Coordinate
	grid map[int]map[int]*Node
}

func NewGrid() *Grid {
	return &Grid{
		grid: make(map[int]map[int]*Node),
	}
}

func (g *Grid) Copy() *Grid {
	newGrid := NewGrid()
	g.Iterate(func(n *Node) {
		newNode := n.Copy()
		newGrid.AddNode(newNode)
	})
	newGrid.data = &bfs.Coordinate{g.data.X, g.data.Y}
	newGrid.free = &bfs.Coordinate{g.free.X, g.free.Y}
	return newGrid
}

func (g *Grid) Move(destination *bfs.Coordinate, source *bfs.Coordinate) {
	if g.free.Equal(destination) {
		g.free = source
	}

	if g.data.Equal(source) {
		g.data = destination
	}

	g.grid[source.Y][source.X].MoveTo(g.grid[destination.Y][destination.X])
}

func (g *Grid) AddNode(n *Node) {
	if _, found := g.grid[n.coordinate.Y]; !found {
		g.grid[n.coordinate.Y] = make(map[int]*Node)
	}
	g.grid[n.coordinate.Y][n.coordinate.X] = n
	if n.used == 0 {
		g.free = n.coordinate
	}

	if n.coordinate.Y == 0 {
		if g.data == nil || n.coordinate.X > g.data.X {
			g.data = n.coordinate
		}
	}
}

func (g *Grid) Equal(node bfs.Node) bool {
	if other, ok := node.(*Grid); ok {
		return g.data.Equal(other.data) && g.free.Equal(other.free)
	}
	return false
}

func (g *Grid) ID() string {
	return g.data.ID() + g.free.ID()
}

func (g *Grid) Neighbors() []bfs.Node {
	neighbors := make([]bfs.Node, 0)
	for _, coordinate := range g.grid[g.free.Y][g.free.X].Neighbors() {
		if row, found := g.grid[coordinate.Y]; found {
			if node, found := row[coordinate.X]; found {
				neighbor := g.Copy()
				neighbor.Move(g.free, coordinate)
				if node.used <= g.grid[g.free.Y][g.free.X].size {
					neighbors = append(neighbors, neighbor)
				}
			}
		}
	}
	return neighbors
}

func (g *Grid) Iterate(cb func(*Node)) {
	for _, row := range g.grid {
		for _, node := range row {
			cb(node)
		}
	}
}

func (g *Grid) String() string {
	/*writer := &bfs.StringWriter{}

	for y := 0; y < len(g.grid); y++ {
		for x := 0; x < len(g.grid[y]); x++ {
			node := g.grid[y][x]
			if node.used == 0 {
				writer.Writef("-%d-", node.size)
			} else {
				writer.Writef("%-3d", node.size)
			}
		}
		writer.Write("\n")
	}

	return writer.String()*/
	return g.ID()
}
