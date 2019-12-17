package main

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	d8 := &D8{}
	challenges[8] = &challenge{"Day 08", "input/day08.txt", d8}
}

type Layer struct {
	width  int
	height int
	rows   [][]int
	index  map[int]int
}

func (l *Layer) parseRow(input []int) error {
	if len(input) != l.width {
		return fmt.Errorf("Expected %d values for image row.  Got %d", l.width, len(input))
	}
	columns := make([]int, len(input))
	for i, v := range input {
		l.index[v] += 1
		columns[i] = v
	}
	l.rows = append(l.rows, columns)
	return nil
}

func (l *Layer) Parse(input []int) (err error) {
	l.index = make(map[int]int)
	for i := 0; i < len(input); i += l.width {
		err = l.parseRow(input[i : i+l.width])
		if err != nil {
			break
		}
	}

	if len(l.rows) < l.height {
		err = fmt.Errorf("Image height %d only have %d rows", l.height, len(l.rows))
	}
	return nil
}

func (l *Layer) Num(digit int) int {
	return l.index[digit]
}

type Image struct {
	width  int
	height int
	layers []*Layer
}

func (img *Image) Parse(input []int) (err error) {
	for i := 0; i < len(input); i += (img.width * img.height) {
		layer := &Layer{width: img.width, height: img.height}
		err = layer.Parse(input[i:(i + (img.width * img.height))])
		if err != nil {
			break
		}
		img.layers = append(img.layers, layer)
	}

	return err
}

func (img *Image) String() string {
	builder := &strings.Builder{}
	for row := 0; row < img.height; row++ {
		for col := 0; col < img.width; col++ {
			for _, l := range img.layers {
				if l.rows[row][col] != 2 {
					if l.rows[row][col] == 0 {
						builder.WriteString(" ")
					} else {
						builder.WriteString("*")
					}
					break
				}
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

type D8 struct {
	input  []int
	width  int
	height int
}

func (d8 *D8) parse(line string) (err error) {
	if len(d8.input) == 0 {
		for _, digit := range strings.Split(line, "") {
			num, err := strconv.Atoi(digit)
			if err == nil {
				d8.input = append(d8.input, num)
			} else {
				return err
			}
		}
	} else {
		err = fmt.Errorf("Image already parsed")
	}
	return err
}

func (d8 *D8) part1() (output string, err error) {
	width := 25
	height := 6
	if d8.width > 0 && d8.height > 0 {
		width = d8.width
		height = d8.height
	}
	img := &Image{width: width, height: height}
	err = img.Parse(d8.input)
	if err == nil {
		layer := img.layers[0]
		count := layer.Num(0)
		for _, l := range img.layers[1:] {
			if l.Num(0) < count {
				count = l.Num(0)
				layer = l
			}
		}

		output = fmt.Sprintf("%d", layer.Num(1)*layer.Num(2))
	}
	return output, err
}

func (d8 *D8) part2() (output string, err error) {
	width := 25
	height := 6
	if d8.width > 0 && d8.height > 0 {
		width = d8.width
		height = d8.height
	}
	img := &Image{width: width, height: height}
	err = img.Parse(d8.input)
	if err == nil {
		output = fmt.Sprintf("\n%s", img.String())
	}
	return output, err
}
