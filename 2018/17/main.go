package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type FillError struct {
	x int
	y int
}

func (fe *FillError) Error() string { return fmt.Sprintf("Full at %d,%d", fe.x, fe.y) }

var (
	ErrOutOfBounds = fmt.Errorf("Exceeded Y boundary")
	ErrClay        = fmt.Errorf("Encountered clay")
	ErrFull        = fmt.Errorf("Tile is full")
)

type Slice struct {
	minX      int
	maxX      int
	maxY      int
	grid      [][]rune
	lastDripX int
	lastDripY int
}

func (slice *Slice) UnmarshalText(input []byte) error {
	coordinates := make(map[int]map[int]struct{})

	for _, line := range bytes.Split(input, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		str := string(line)
		x1, x2, y1, y2 := 0, 0, 0, 0
		if _, err := fmt.Sscanf(str, "x=%d, y=%d..%d", &x1, &y1, &y2); err == nil {
			x2 = x1
		} else if _, err := fmt.Sscanf(str, "y=%d, x=%d..%d", &y1, &x1, &x2); err == nil {
			y2 = y1
		} else if _, err := fmt.Sscanf(str, "x=%d, y=%d", &x1, &y1); err == nil {
			x2 = x1
			y2 = y1
		} else {
			return fmt.Errorf("Unrecognized line: %q\n", str)
		}

		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				if len(coordinates) == 0 {
					slice.minX = x - 1
					slice.maxX = x

					slice.maxY = y
				} else {
					if x < slice.minX {
						slice.minX = x - 1
					}

					if x > slice.maxX {
						slice.maxX = x
					}

					if y > slice.maxY {
						slice.maxY = y
					}
				}
				if row, found := coordinates[y]; found {
					row[x] = struct{}{}
				} else {
					coordinates[y] = make(map[int]struct{})
					coordinates[y][x] = struct{}{}
				}
			}
		}
	}

	if slice.maxX < 500 {
		slice.maxX = 500
	}

	slice.grid = make([][]rune, slice.maxY+1)
	for y := 0; y <= slice.maxY; y++ {
		slice.grid[y] = make([]rune, slice.maxX-slice.minX+2)
		for x := slice.minX; x <= slice.maxX+1; x++ {
			if _, found := coordinates[y][x]; found {
				slice.grid[y][x-slice.minX] = '#'
			} else {
				slice.grid[y][x-slice.minX] = '.'
			}
		}
	}
	slice.grid[0][500-slice.minX] = '+'
	return nil
}

func (slice *Slice) Fill() {
	slice.lastDripX = 500 - slice.minX
	slice.lastDripY = 0
	slice.fill(500-slice.minX, 1)
}

func (slice *Slice) fill(x, y int) (err error) {
	if y < 0 || x < 0 || len(slice.grid) <= y || len(slice.grid[0]) <= x {
		return
	}

	if slice.grid[y][x] == '#' {
		return
	}

	if slice.grid[y][x] == '.' {
		slice.grid[y][x] = '|'
		if len(slice.grid) <= y+1 {
			return
		}

		slice.fill(x, y+1)

		if slice.grid[y+1][x] == '~' || slice.grid[y+1][x] == '#' {
			// get bounds if any
			leftX, rightX := x, x
			leftBound, rightBound := false, false
			for ; 0 <= leftX; leftX-- {
				if slice.grid[y+1][leftX] == '.' {
					break
				}

				if slice.grid[y][leftX] == '#' {
					leftBound = true
					break
				}
			}
			leftX++

			for ; rightX < len(slice.grid[y]); rightX++ {
				if slice.grid[y+1][rightX] == '.' {
					break
				}

				if slice.grid[y][rightX] == '#' {
					rightBound = true
					break
				}
			}

			fill := '|'
			if leftBound && rightBound {
				fill = '~'
			}

			for xx := leftX; xx < rightX; xx++ {
				slice.grid[y][xx] = fill
			}

			if !leftBound {
				slice.fill(leftX-1, y)
			}

			if !rightBound {
				slice.fill(rightX, y)
			}
		}
	}
	return
}

func (slice *Slice) Filled() int {
	count := 0
	for _, row := range slice.grid[0 : len(slice.grid)-1] {
		for _, tile := range row {
			if tile == '|' || tile == '~' {
				count++
			}
		}
	}
	return count
}

func (slice *Slice) AtRest() int {
	count := 0
	for _, row := range slice.grid[0 : len(slice.grid)-1] {
		for _, tile := range row {
			if tile == '~' {
				count++
			}
		}
	}
	return count
}

func (slice *Slice) String() string {
	runes := []rune{}
	for _, row := range slice.grid {
		runes = append(runes, row...)
		runes = append(runes, '\n')
	}
	return string(runes)
}

func part1(input []byte) error {
	slice := &Slice{}
	err := slice.UnmarshalText(input)
	if err == nil {
		slice.Fill()
		filled := slice.Filled()
		fmt.Printf("Part 1: %d\n", filled)
	}
	return err
}

func part2(input []byte) error {
	slice := &Slice{}
	err := slice.UnmarshalText(input)
	if err == nil {
		slice.Fill()
		atRest := slice.AtRest()
		fmt.Printf("Part 2: %d\n", atRest)
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
		err = f(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Part %d failed: %v\n", i+1, err)
			os.Exit(-1)
		}
	}
}
