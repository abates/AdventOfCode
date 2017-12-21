package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Transforms struct {
	transforms map[string][]string
}

func NewTransforms() *Transforms {
	return &Transforms{
		transforms: make(map[string][]string),
	}
}

func (t *Transforms) Add(definition string) {
	if definition == "" {
		return
	}
	fields := strings.Split(definition, " => ")
	match := strings.Replace(fields[0], "/", "", -1)
	replace := strings.Split(fields[1], "/")
	t.transforms[match] = replace
}

func (t *Transforms) Transform(block *Block) *Block {
	for i := 0; i < 4; i++ {
		if newBlock, found := t.transforms[block.String()]; found {
			return NewBlock(newBlock)
		}

		tmpBlock := block.FlipVertical()
		if newBlock, found := t.transforms[tmpBlock.String()]; found {
			return NewBlock(newBlock)
		}

		tmpBlock = block.FlipHorizontal()
		if newBlock, found := t.transforms[tmpBlock.String()]; found {
			return NewBlock(newBlock)
		}

		block = block.Rotate()
	}
	panic("Can't transform")
}

type Block struct {
	rows []string
}

func NewBlock(rows []string) *Block {
	block := &Block{
		rows: make([]string, len(rows)),
	}

	copy(block.rows, rows)
	return block
}

func (b Block) FlipHorizontal() *Block {
	rows := make([][]byte, len(b.rows))
	for i, _ := range rows {
		row := make([]byte, len([]byte(b.rows[i])))
		rows[i] = row
		for j, c := range []byte(b.rows[i]) {
			row[len(b.rows[i])-j-1] = c
		}
	}

	str := make([]string, len(rows))
	for i, row := range rows {
		str[i] = string(row)
	}
	return &Block{str}
}

func (b Block) FlipVertical() *Block {
	rows := make([]string, len(b.rows))
	for i := 0; i < len(rows); i++ {
		rows[i] = b.rows[len(rows)-i-1]
	}
	return &Block{rows}
}

func (b Block) Rotate() *Block {
	output := [][]byte{}
	for _, row := range b.rows {
		output = append(output, []byte(row))
	}

	n := len(output)
	for x := 0; x < n/2; x++ {
		for y := x; y < n-x-1; y++ {
			temp := output[x][y]
			output[x][y] = output[y][n-1-x]
			output[y][n-1-x] = output[n-1-x][n-1-y]
			output[n-1-x][n-1-y] = output[n-1-y][x]
			output[n-1-y][x] = temp
		}
	}

	rows := make([]string, n)
	for i, b := range output {
		rows[i] = string(b)
	}
	return &Block{rows}
}

func (b Block) String() string { return strings.Join(b.rows, "") }

type Grid struct {
	pixels []string
}

func NewGrid(input []string) *Grid {
	return &Grid{input}
}

func (g *Grid) Divide(transforms *Transforms) {
	sliceSize := 3
	newSize := 4
	if len(g.pixels)%2 == 0 {
		sliceSize = 2
		newSize = 3
	}

	newPixels := make([]bytes.Buffer, 0)
	for row := 0; row < len(g.pixels); row += sliceSize {
		pixelRows := make([]bytes.Buffer, newSize)
		for column := 0; column < len(g.pixels[row]); column += sliceSize {
			block := make([]string, sliceSize)
			for i := 0; i < sliceSize; i++ {
				block[i] = g.pixels[row+i][column : column+sliceSize]
			}

			for i, blockRow := range transforms.Transform(NewBlock(block)).rows {
				pixelRows[i].WriteString(blockRow)
			}
		}
		newPixels = append(newPixels, pixelRows...)
	}

	for i := 0; i < len(newPixels); i++ {
		if i < len(g.pixels) {
			g.pixels[i] = newPixels[i].String()
		} else {
			g.pixels = append(g.pixels, newPixels[i].String())
		}
	}
}

func (g *Grid) NumPixelsOn() int {
	num := 0
	for _, row := range g.pixels {
		num += strings.Count(row, "#")
	}
	return num
}

func (g *Grid) String() string {
	return strings.Join(g.pixels, "\n")
}

func StandardGrid() *Grid {
	return NewGrid([]string{".#.", "..#", "###"})
}

func main() {
	transforms := NewTransforms()
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	for _, line := range strings.Split(string(b), "\n") {
		transforms.Add(strings.TrimSpace(line))
	}

	grid := StandardGrid()
	for i := 0; i < 5; i++ {
		grid.Divide(transforms)
	}

	fmt.Printf("Part 1: %d\n", grid.NumPixelsOn())

	for i := 0; i < 13; i++ {
		grid.Divide(transforms)
	}
	fmt.Printf("Part 2: %d\n", grid.NumPixelsOn())
}
