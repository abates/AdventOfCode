package main

import (
	"fmt"
)

type Buffer struct {
	values []int
	pos    int
}

func (b *Buffer) Spin(steps int) {
	if len(b.values) > 0 {
		b.pos = (b.pos + steps) % len(b.values)
	}
}

func (b *Buffer) Insert(value int) {
	if len(b.values) == 0 {
		b.values = append(b.values, value)
	} else {
		b.values = append(b.values[0:b.pos+1], append([]int{value}, b.values[b.pos+1:]...)...)
	}
	b.Spin(1)
}

func valueAfter(steps, iterations, after int) int {
	last := 0
	pos := 0
	for i := 1; i < iterations+1; i++ {
		pos = (pos+steps)%i + 1
		if pos == after+1 {
			last = i
		}
	}
	return last
}

func main() {
	//run(3, 10)
	step := 359
	part1Iterations := 2017
	part2Iterations := 50000000

	b := &Buffer{}
	for i := 0; i < part1Iterations; i++ {
		b.Insert(i)
		b.Spin(step)
	}
	b.Spin(1)
	fmt.Printf("Part 1: %d\n", b.values[b.pos])

	value := valueAfter(step, part2Iterations, 0)
	fmt.Printf("Part 2: %d\n", value)
}
