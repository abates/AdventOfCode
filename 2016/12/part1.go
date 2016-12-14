package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
	"strings"
)

func main() {
	program := make([][]string, 0)
	for _, line := range util.ReadInput() {
		program = append(program, strings.Split(line, " "))
	}
	part1(program)
	part2(program)
}

func part1(program [][]string) {
	c := &Computer{
		registers: make(map[string]int),
	}

	c.Run(program)

	fmt.Printf("Computer state:\n%s\n", c)
}

func part2(program [][]string) {
	c := &Computer{
		registers: make(map[string]int),
	}
	c.registers["c"] = 1

	c.Run(program)

	fmt.Printf("Computer state:\n%s\n", c)
}
