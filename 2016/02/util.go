package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
	"strings"
)

func readSequences() [][]string {
	sequences := make([][]string, 0)
	for _, line := range util.ReadInput() {
		sequences = append(sequences, strings.Split(line, ""))
	}
	return sequences
}

type Keypad struct {
	x    int
	y    int
	keys [][]string
}

func NewKeypad(startx, starty int, keys [][]string) *Keypad {
	if keys == nil {
		keys = [][]string{
			[]string{"1", "2", "3"},
			[]string{"4", "5", "6"},
			[]string{"7", "8", "9"},
		}
	}

	k := &Keypad{
		x:    startx,
		y:    starty,
		keys: keys,
	}
	return k
}

func (k *Keypad) move(x, y int) {
	if x != 0 && k.x+x >= 0 && k.x+x < len(k.keys[y]) && k.keys[k.y][k.x+x] != "0" {
		k.x += x
	}

	if y != 0 && k.y+y >= 0 && k.y+y < len(k.keys) && k.keys[k.y+y][k.x] != "0" {
		k.y += y
	}
}

func (k *Keypad) Move(direction string) (err error) {
	switch direction {
	case "U":
		k.move(0, -1)
	case "R":
		k.move(1, 0)
	case "D":
		k.move(0, 1)
	case "L":
		k.move(-1, 0)
	default:
		err = fmt.Errorf("don't know how to move %s", direction)
	}
	return
}

func (k *Keypad) Position() string {
	return k.keys[k.y][k.x]
}
