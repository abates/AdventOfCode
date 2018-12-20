package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type Landscape struct {
	layout [][]rune
}

func (land *Landscape) UnmarshalText(text []byte) error {
	land.layout = nil
	for _, line := range bytes.Split(text, []byte("\n")) {
		land.layout = append(land.layout, []rune(string(line)))
	}
	return nil
}

func (land *Landscape) Advance(minutes int) {
	repeating := false
	index := make(map[string]int)
	index[land.String()] = 0

	for count := 0; count < minutes; count++ {
		land.advance()
		if c, found := index[land.String()]; !repeating && found {
			period := count - c
			minutes = count + (minutes-count)%period
			repeating = true
		}

		if !repeating {
			index[land.String()] = count
		}
	}
	return
}

func (land *Landscape) String() string {
	runes := []rune{}
	for _, row := range land.layout {
		runes = append(runes, row...)
		runes = append(runes, '\n')
	}
	return string(runes)
}

func (land *Landscape) countAdjacent(x, y int) (trees, lumberyards int) {
	for _, delta := range [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}} {
		deltaX, deltaY := delta[0], delta[1]
		if y+deltaY < 0 || len(land.layout) <= y+deltaY {
			continue
		}
		row := land.layout[y+deltaY]
		if x+deltaX < 0 || len(row) <= x+deltaX {
			continue
		}
		r := row[x+deltaX]
		if r == '|' {
			trees++
		} else if r == '#' {
			lumberyards++
		}
	}
	return
}

type Change struct {
	x int
	y int
	r rune
}

func (land *Landscape) advance() {
	changes := []Change{}
	for y, row := range land.layout {
		for x, r := range row {
			trees, lumberyards := land.countAdjacent(x, y)
			if r == '.' {
				if trees > 2 {
					changes = append(changes, Change{x, y, '|'})
				}
			} else if r == '|' {
				if lumberyards > 2 {
					changes = append(changes, Change{x, y, '#'})
				}
			} else if r == '#' {
				if lumberyards == 0 || trees == 0 {
					changes = append(changes, Change{x, y, '.'})
				}
			}
		}
	}
	for _, change := range changes {
		land.layout[change.y][change.x] = change.r
	}
}

func (land *Landscape) count(match rune) int {
	count := 0
	for _, row := range land.layout {
		for _, r := range row {
			if r == match {
				count++
			}
		}
	}
	return count
}

func (land *Landscape) Wood() int {
	return land.count('|')
}

func (land *Landscape) Lumber() int {
	return land.count('#')
}

func part1(input []byte) error {
	land := &Landscape{}
	err := land.UnmarshalText(input)
	if err == nil {
		land.Advance(10)
		wood := land.Wood()
		lumber := land.Lumber()
		fmt.Printf("Part 1: %d\n", wood*lumber)
	}
	return err
}

func part2(input []byte) error {
	land := &Landscape{}
	err := land.UnmarshalText(input)
	if err == nil {
		land.Advance(1000000000)
		wood := land.Wood()
		lumber := land.Lumber()
		fmt.Printf("Part 2: %d\n", wood*lumber)
	}
	return err
}

func main() {
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
			fmt.Fprintf(os.Stderr, "Part %d failed: %v\n", i, err)
			os.Exit(-1)
		}
	}
}
