package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/abates/AdventOfCode/coordinate"
)

func init() {
	d11 := &D11{}
	challenges[11] = &challenge{"Day 11", "input/day11.txt", d11}
}

type direction int

var directions = []coordinate.Coordinate{
	coordinate.New(0, 1),
	coordinate.New(1, 0),
	coordinate.New(0, -1),
	coordinate.New(-1, 0),
}

func (d *direction) Update(str string) coordinate.Coordinate {
	if str == "0" {
		*d -= 1
		if *d < 0 {
			*d = 3
		}
	} else if str == "1" {
		*d += 1
		if *d > 3 {
			*d = 0
		}
	} else {
		panic(fmt.Sprintf("Unknown direction %s", str))
	}
	return directions[*d]
}

type robot struct {
	pos        coordinate.Coordinate
	direction  direction
	mem        []*Int
	panels     map[float64]map[float64]bool
	numPainted int
	minX       float64
	minY       float64
	maxX       float64
	maxY       float64
}

func (r *robot) Get(x, y float64) (bool, bool) {
	row := r.panels[y]
	if row == nil {
		return false, false
	}
	v, found := row[x]
	return v, found
}

func (r *robot) Set(x, y float64, value bool) {
	if x < r.minX {
		r.minX = x
	}

	if r.maxX < x {
		r.maxX = x
	}

	if y < r.minY {
		r.minY = y
	}

	if r.maxY < y {
		r.maxY = y
	}

	row := r.panels[y]
	if row == nil {
		row = make(map[float64]bool)
		r.panels[y] = row
	}
	row[x] = value
}

func (r *robot) Dump() string {
	builder := &strings.Builder{}
	for y := r.maxY; y >= r.minY; y-- {
		for x := r.minX; x <= r.maxX; x++ {
			if v, _ := r.Get(x, y); v {
				builder.WriteString("#")
			} else {
				builder.WriteString(" ")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (r *robot) Run(initial bool) {
	r.panels = make(map[float64]map[float64]bool)
	r.Set(0, 0, initial)
	pr, input := io.Pipe()
	output, pw := io.Pipe()
	computer := NewComputer(r.mem)
	computer.SetInput(pr)
	computer.SetOutput(pw)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		computer.Run()
		input.Close()
		pw.Close()
		wg.Done()
	}()

	rr := bufio.NewReader(output)
	//ww := bufio.NewWriter(input)
	var err error
	var line []byte
	for err == nil {
		if v, _ := r.Get(r.pos.Get(0), r.pos.Get(1)); v {
			fmt.Fprintf(input, "1\n")
		} else {
			fmt.Fprintf(input, "0\n")
		}

		line, _, err = rr.ReadLine()
		if err == nil {
			if _, found := r.Get(r.pos.Get(0), r.pos.Get(1)); !found {
				r.numPainted++
			}

			if string(line) == "0" {
				r.Set(r.pos.Get(0), r.pos.Get(1), false)
			} else if string(line) == "1" {
				r.Set(r.pos.Get(0), r.pos.Get(1), true)
			} else {
				panic(fmt.Sprintf("Unknown color instruction, %q", line))
			}

			line, _, err = rr.ReadLine()
			if err == nil {
				delta := r.direction.Update(string(line))
				r.pos = r.pos.Add(delta)
			}
		}
	}
	wg.Wait()
}

type D11 struct {
	mem []*Int
}

func (d11 *D11) parseFile(lines []string) (err error) {
	d11.mem, err = ParseComputerMemory(lines)
	return err
}

func (d11 *D11) part1() (string, error) {
	r := &robot{mem: d11.mem, pos: coordinate.New(0, 0)}
	r.Run(false)
	return fmt.Sprintf("%d panels painted", r.numPainted), nil
}

func (d11 *D11) part2() (string, error) {
	r := &robot{mem: d11.mem, pos: coordinate.New(0, 0)}
	r.Run(true)
	return fmt.Sprintf("\n%s", r.Dump()), nil
}
