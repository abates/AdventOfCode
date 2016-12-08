package main

import (
	"strings"
)

type Screen struct {
	pixels [][]string
}

func NewScreen(width, height int) *Screen {
	pixels := make([][]string, height)
	for i, _ := range pixels {
		for j := 0; j < width; j++ {
			pixels[i] = append(pixels[i], " ")
		}
	}

	return &Screen{
		pixels: pixels,
	}
}

func (s *Screen) String() string {
	lines := make([]string, 0)
	for _, row := range s.pixels {
		lines = append(lines, strings.Join(row, ""))
	}
	return strings.Join(lines, "\n")
}

func (s *Screen) Rect(width, height int) {
	for row := 0; row < height; row++ {
		if row >= len(s.pixels) {
			break
		}
		for column := 0; column < width; column++ {
			if column >= len(s.pixels[row]) {
				break
			}
			s.pixels[row][column] = "#"
		}
	}
}

func rotate(line []string, amount int) []string {
	return append(line[len(line)-amount:], line[:len(line)-amount]...)
}

func (s *Screen) RotateRow(row, amount int) {
	if row >= len(s.pixels) {
		return
	}

	s.pixels[row] = rotate(s.pixels[row], amount)
}

func (s *Screen) RotateColumn(column, amount int) {
	values := make([]string, len(s.pixels))
	for i, row := range s.pixels {
		if column >= len(row) {
			return
		}
		values[i] = row[column]
	}

	values = rotate(values, amount)

	for i, row := range s.pixels {
		row[column] = values[i]
	}
}
