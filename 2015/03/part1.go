package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func updatePresents(presents map[string]int, x, y int) {
	key := fmt.Sprintf("(%d,%d)", x, y)
	if _, found := presents[key]; found {
		presents[key]++
	} else {
		presents[key] = 1
	}
}

func main() {
	presents := make(map[string]int)
	file, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(file)
	x := 0
	y := 0
	updatePresents(presents, x, y)
	for _, c := range string(b) {
		switch c {
		case '^':
			y++
		case 'v':
			y--
		case '<':
			x--
		case '>':
			x++
		}
		updatePresents(presents, x, y)
	}
	fmt.Printf("%d\n", len(presents))
}
