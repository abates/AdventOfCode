package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/abates/AdventOfCode/coordinate"
)

var (
	Up    = coordinate.New(0, -1)
	Down  = coordinate.New(0, 1)
	Right = coordinate.New(1, 0)
	Left  = coordinate.New(-1, 0)

	Directions = map[string]*coordinate.Coordinate{"UP": Up, "Down": Down, "Right": Right, "Left": Left}
)

type Diagram struct {
	grid map[int]map[int]string
}

func (d *Diagram) Lookup(c *coordinate.Coordinate) string {
	if row, found := d.grid[c.Coordinates[1]]; found {
		return row[c.Coordinates[0]]
	}
	return ""
}

func (d *Diagram) Set(x, y int, s string) {
	if d.grid == nil {
		d.grid = make(map[int]map[int]string)
	}

	if d.grid[y] == nil {
		d.grid[y] = make(map[int]string)
	}

	d.grid[y][x] = s
}

type Packet struct {
	prevDir *coordinate.Coordinate
	prevPos *coordinate.Coordinate
	curPos  *coordinate.Coordinate
	diagram *Diagram
	path    []string
	steps   int
}

func (p *Packet) move(nextPos *coordinate.Coordinate) {
	p.prevPos = p.curPos
	p.curPos = nextPos
}

func (p *Packet) Move() bool {
	p.steps++
	p.move(p.curPos.Add(p.prevDir))
	current := p.diagram.Lookup(p.curPos)
	if current == "" {
		return false
	}

	if current == "+" {
		for _, direction := range Directions {
			nextPos := p.curPos.Add(direction)
			n := p.diagram.Lookup(nextPos)
			if n != "" && !nextPos.Equal(p.prevPos) {
				p.prevDir = direction
				break
			}
		}
	} else if strings.IndexAny(current, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") != -1 {
		p.path = append(p.path, current)
	}

	return true
}

func (p *Packet) Traverse() {
	startX := 0
	for i, s := range p.diagram.grid[0] {
		if s != " " {
			startX = i
		}
	}

	p.curPos = coordinate.New(startX, 0)
	for p.Move() {
	}
}

func parseInput(input string) *Diagram {
	diagram := &Diagram{}

	for y, line := range strings.Split(input, "\n") {
		for x, s := range strings.Split(line, "") {
			if s != " " {
				diagram.Set(x, y, s)
			}
		}
	}
	return diagram
}

func main() {
	input := `     |
     |  +--+    
     A  |  C    
 F---|--|-E---+ 
     |  |  |  D 
     +B-+  +--+ 
	`
	packet := &Packet{diagram: parseInput(input), prevDir: Down}
	packet.Traverse()
	fmt.Printf("Part 1: %s\n", strings.Join(packet.path, ""))
	fmt.Printf("Part 2: %d\n", packet.steps)

	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	input = string(b)
	packet = &Packet{diagram: parseInput(input), prevDir: Down}
	packet.Traverse()

	fmt.Printf("Part 1: %s\n", strings.Join(packet.path, ""))
	fmt.Printf("Part 2: %d\n", packet.steps)
}
