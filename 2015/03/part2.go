package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func updatePresents(presents map[string]int, m *mover) {
	key := fmt.Sprintf("(%d,%d)", m.x, m.y)
	if _, found := presents[key]; found {
		presents[key]++
	} else {
		presents[key] = 1
	}
}

type mover struct {
	x int
	y int
}

func (m *mover) move(direction rune) {
	switch direction {
	case '^':
		m.y++
	case 'v':
		m.y--
	case '<':
		m.x--
	case '>':
		m.x++
	}
}

func main() {
	presents := make(map[string]int)
	file, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(file)
	santa := &mover{}
	robot := &mover{}
	updatePresents(presents, santa)
	for i, c := range string(b) {
		var m *mover
		if i%2 == 0 {
			m = santa
		} else {
			m = robot
		}
		m.move(c)
		updatePresents(presents, m)
	}
	fmt.Printf("%d\n", len(presents))
}
