package main

import (
	"fmt"
	"os"
	"strconv"
)

type Grid struct {
	serialNumber int
	digit        int
	size         int
	index        map[int]map[int]map[int]int
}

func (g *Grid) setGridLevel(x, y, size, level int) {
	if g.index == nil {
		g.index = make(map[int]map[int]map[int]int)
	}
	if _, found := g.index[x]; !found {
		g.index[x] = make(map[int]map[int]int)
	}
	if _, found := g.index[x][y]; !found {
		g.index[x][y] = make(map[int]int)
	}
	g.index[x][y][size] = level
}

func (g *Grid) getGridLevel(x, y, size int) (int, bool) {
	if g.index == nil {
		return 0, false
	}
	if col, found := g.index[x]; found {
		if index, found := col[y]; found {
			if level, found := index[size]; found {
				return level, true
			}
		}
	}
	return 0, false
}

func (g *Grid) FuelLevel(x, y int) int {
	rackId := x + 10
	power := rackId*y + g.serialNumber
	power *= rackId
	power = (power / g.digit) % 10
	power -= 5
	return power
}

func (g *Grid) GridLevel(x, y, size int) int {
	level := 0
	found := false
	if level, found = g.getGridLevel(x, y, size-1); found {
		// get last column
		for cx := x; cx < x+size; cx++ {
			level += g.FuelLevel(cx, y+size-1)
		}

		// get last row
		for cy := y; cy < y+size-1; cy++ {
			level += g.FuelLevel(x+size-1, cy)
		}
	} else {
		for cx := x; cx < x+size; cx++ {
			for cy := y; cy < y+size; cy++ {
				level += g.FuelLevel(cx, cy)
			}
		}
	}
	g.setGridLevel(x, y, size, level)
	return level
}

func (g *Grid) MaxFuelLevel(gridSize int) (level, x, y int) {
	x = 1
	y = 1
	level = g.GridLevel(x, y, gridSize)
	for cx := 1; cx <= g.size-gridSize; cx++ {
		for cy := 1; cy <= g.size-gridSize; cy++ {
			clevel := g.GridLevel(cx, cy, gridSize)
			if clevel > level {
				x = cx
				y = cy
				level = clevel
			}
		}
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <serial number>\n", os.Args[0])
		os.Exit(-1)
	}
	serialNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "serial number must be an integer: %v\n", err)
		os.Exit(-1)
	}
	g := &Grid{
		serialNumber: serialNumber,
		digit:        100,
		size:         300,
	}

	maxLevel, maxX, maxY := g.MaxFuelLevel(3)
	fmt.Printf("Max (3x3) Fuel %d at (%d, %d)\n", maxLevel, maxX, maxY)

	maxSize := 0
	for size := 3; size <= g.size; size++ {
		level, x, y := g.MaxFuelLevel(size)
		if level > maxLevel || size == 3 {
			maxLevel = level
			maxX = x
			maxY = y
			maxSize = size
		}
	}
	fmt.Printf("Max (%dx%d) Fuel %d at (%d, %d)\n", maxSize, maxSize, maxLevel, maxX, maxY)
}
