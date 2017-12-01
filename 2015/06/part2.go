package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Row struct {
	lights []int
}

func (r *Row) TurnOn(x1, x2 int) {
	for x := x1; x <= x2; x++ {
		r.lights[x]++
	}
}

func (r *Row) TurnOff(x1, x2 int) {
	for x := x1; x <= x2; x++ {
		r.lights[x]--
		if r.lights[x] < 0 {
			r.lights[x] = 0
		}
	}
}

func (r *Row) Toggle(x1, x2 int) {
	for x := x1; x <= x2; x++ {
		r.lights[x] += 2
	}
}

func (r *Row) Brightness() int {
	brightness := 0
	for _, light := range r.lights {
		brightness += light
	}
	return brightness
}

type Grid struct {
	rows []*Row
}

func NewGrid() *Grid {
	grid := &Grid{
		rows: make([]*Row, 1000),
	}

	for i, _ := range grid.rows {
		grid.rows[i] = &Row{
			lights: make([]int, 1000),
		}
	}
	return grid
}

func (g *Grid) TurnOn(x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		g.rows[y].TurnOn(x1, x2)
	}
}

func (g *Grid) TurnOff(x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		g.rows[y].TurnOff(x1, x2)
	}
}

func (g *Grid) Toggle(x1, y1, x2, y2 int) {
	for y := y1; y <= y2; y++ {
		g.rows[y].Toggle(x1, x2)
	}
}

func (g *Grid) Brightness() int {
	brightness := 0
	for _, row := range g.rows {
		brightness += row.Brightness()
	}
	return brightness
}

func main() {
	grid := NewGrid()

	f, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(f)
	for _, line := range strings.Split(string(b), "\n") {
		var x1, x2, y1, y2 int
		if strings.HasPrefix(line, "turn on") {
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			grid.TurnOn(x1, y1, x2, y2)
		} else if strings.HasPrefix(line, "turn off") {
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			grid.TurnOff(x1, y1, x2, y2)
		} else if strings.HasPrefix(line, "toggle") {
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &x1, &y1, &x2, &y2)
			grid.Toggle(x1, y1, x2, y2)
		}
	}

	fmt.Printf("Total Brightness: %d\n", grid.Brightness())
}
