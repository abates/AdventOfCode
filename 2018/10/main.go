package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Point struct {
	x      int
	y      int
	deltaX int
	deltaY int
}

func (p *Point) UnmarshalText(text []byte) error {
	str := string(text)
	_, err := fmt.Sscanf(str, "position=<%d,  %d> velocity=<%d, %d>", &p.x, &p.y, &p.deltaX, &p.deltaY)
	return err
}

func (p *Point) advance() {
	p.x += p.deltaX
	p.y += p.deltaY
}

type Grid struct {
	minX       int
	minY       int
	maxX       int
	maxY       int
	points     []*Point
	pointIndex map[int]map[int]*Point
}

func (g *Grid) AddText(text []byte) error {
	point := &Point{}
	err := point.UnmarshalText(text)
	if err == nil {
		if len(g.points) == 0 {
			g.minX = point.x
			g.minY = point.y
			g.maxX = point.x
			g.maxY = point.y
		}
		g.points = append(g.points, point)
		g.index(point)
	}

	return err
}

func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}

func max(i1, i2 int) int {
	if i2 < i1 {
		return i1
	}
	return i2
}

func (g *Grid) index(point *Point) {
	g.minX = min(g.minX, point.x)
	g.minY = min(g.minY, point.y)
	g.maxX = max(g.maxX, point.x)
	g.maxY = max(g.maxY, point.y)

	if g.pointIndex == nil {
		g.pointIndex = make(map[int]map[int]*Point)
	}

	if row, found := g.pointIndex[point.x]; found {
		row[point.y] = point
	} else {
		g.pointIndex[point.x] = make(map[int]*Point)
		g.pointIndex[point.x][point.y] = point
	}
}

func (g *Grid) advance() {
	g.pointIndex = nil
	g.points[0].advance()
	g.minX = g.points[0].x
	g.minY = g.points[0].y
	g.maxX = g.points[0].x
	g.maxY = g.points[0].y

	for _, point := range g.points[1:] {
		point.advance()
		g.index(point)
	}
}

func (g *Grid) Width() int {
	return g.maxX - g.minX
}

func (g *Grid) Height() int {
	return g.maxY - g.minY
}

func (g *Grid) String() string {
	var builder strings.Builder
	for y := g.minY; y <= g.maxY; y++ {
		fields := []string{}
		for x := g.minX; x <= g.maxX; x++ {
			if _, found := g.pointIndex[x][y]; found {
				fields = append(fields, "#")
			} else {
				fields = append(fields, ".")
			}
		}
		builder.WriteString(fmt.Sprintf("%s\n", strings.Join(fields, "")))
	}
	return builder.String()
}

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Error: " + err.Error())
	}

	grid := &Grid{}

	for _, line := range bytes.Split(input, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		err := grid.AddText(line)
		if err != nil {
			panic("Error: " + err.Error())
		}
	}

	for second := 1; ; second++ {
		grid.advance()
		if grid.Width() < 80 && grid.Height() < 25 {
			fmt.Printf("Second %d\n", second)
			fmt.Printf("%s\n", grid.String())
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
}
