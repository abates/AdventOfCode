package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/abates/AdventOfCode/coordinate"
)

type Direction int

const (
	Left  = Direction(-1)
	Right = Direction(1)
)

var Directions = []*coordinate.Coordinate{
	coordinate.New(0, -1),
	coordinate.New(1, 0),
	coordinate.New(0, 1),
	coordinate.New(-1, 0),
}

type Carrier struct {
	position    *coordinate.Coordinate
	direction   Direction
	NumInfected int
	grid        *Grid
}

func NewCarrier(grid *Grid) *Carrier {
	return &Carrier{
		position: coordinate.New(grid.Width()/2, grid.Height()/2),
		grid:     grid,
	}
}

func (c *Carrier) Burst() {
	node := c.grid.Lookup(c.position)
	if node.Infected() {
		node.Clean()
		c.Turn(Right)
	} else {
		node.Infect()
		c.Turn(Left)
		c.NumInfected++
	}
	c.Move()
}

func (c *Carrier) Turn(direction Direction) {
	c.direction += direction
	if c.direction < 0 {
		c.direction = Direction(len(Directions) - 1)
	} else if int(c.direction) == len(Directions) {
		c.direction = 0
	}
}

func (c *Carrier) Move() {
	c.position = c.position.Add(Directions[c.direction])
}

type Node struct {
	infected bool
}

func (n *Node) Clean() {
	n.infected = false
}

func (n *Node) Infect() {
	n.infected = true
}

func (n *Node) Infected() bool {
	return n.infected
}

type Grid struct {
	nodes map[int]map[int]*Node
	maxX  int
	minX  int
	maxY  int
	minY  int
}

func NewGrid() *Grid {
	return &Grid{
		nodes: make(map[int]map[int]*Node),
	}
}

func (g *Grid) Set(position *coordinate.Coordinate, node *Node) {
	column, found := g.nodes[position.Coordinates[0]]
	if !found {
		column = make(map[int]*Node)
		g.nodes[position.Coordinates[0]] = column
	}

	column[position.Coordinates[1]] = node

	if position.Coordinates[0] > g.maxX {
		g.maxX = position.Coordinates[0]
	}
	if position.Coordinates[0] < g.minX {
		g.minX = position.Coordinates[0]
	}

	if position.Coordinates[1] > g.maxY {
		g.maxY = position.Coordinates[1]
	}
	if position.Coordinates[1] < g.minY {
		g.minY = position.Coordinates[1]
	}
}

func (g *Grid) Width() int {
	return g.maxX - g.minX
}

func (g *Grid) Height() int {
	return g.maxY - g.minY
}

func (g *Grid) Lookup(position *coordinate.Coordinate) *Node {
	column, found := g.nodes[position.Coordinates[0]]
	if !found {
		column := make(map[int]*Node)
		g.nodes[position.Coordinates[0]] = column
	}

	node, found := column[position.Coordinates[1]]
	if !found {
		node = &Node{}
		g.Set(coordinate.New(position.Coordinates[0], position.Coordinates[1]), node)
	}
	return node
}

func ParseGrid(input string) *Grid {
	g := NewGrid()
	for y, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, value := range strings.Split(strings.TrimSpace(line), "") {
			node := &Node{}
			if value == "#" {
				node.Infect()
			}
			g.Set(coordinate.New(x, y), node)
		}
	}
	return g
}

func main() {
	/*input := "..#\n#..\n..."

	grid := ParseGrid(input)
	carrier := NewCarrier(grid)
	for i := 0; i < 10000; i++ {
		carrier.Burst()
	}

	fmt.Printf("Test: %d infected\n", carrier.NumInfected)*/

	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	grid := ParseGrid(string(b))
	carrier := NewCarrier(grid)

	for i := 0; i < 10000; i++ {
		carrier.Burst()
	}

	fmt.Printf("Part 1: %d infected\n", carrier.NumInfected)
}
