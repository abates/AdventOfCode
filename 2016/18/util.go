package main

import (
	"strings"
)

type Tile string

func (t Tile) IsSafe() bool {
	return string(t) == "."
}

func (t Tile) IsTrap() bool {
	return string(t) == "^"
}

func NewTile(left, center, right Tile) Tile {
	if (left.IsTrap() && right.IsSafe()) || (left.IsSafe() && right.IsTrap()) {
		return Tile("^")
	}

	return Tile(".")
}

type Row struct {
	tiles []Tile
}

func NewRowFromString(s string) *Row {
	row := &Row{}
	for _, tile := range strings.Split(s, "") {
		row.tiles = append(row.tiles, Tile(tile))
	}
	return row
}

func (r *Row) NextRow() *Row {
	nextRow := &Row{}

	var left, center, right Tile
	for i := 0; i < len(r.tiles); i++ {
		if i == 0 {
			left = Tile(".")
		} else {
			left = r.tiles[i-1]
		}

		center = r.tiles[i]

		if i == len(r.tiles)-1 {
			right = Tile(".")
		} else {
			right = r.tiles[i+1]
		}

		nextRow.tiles = append(nextRow.tiles, NewTile(left, center, right))
	}
	return nextRow
}

func (r *Row) NumSafe() int {
	safe := 0
	for _, tile := range r.tiles {
		if tile.IsSafe() {
			safe++
		}
	}
	return safe
}
