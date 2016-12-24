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
	//part1(program)
	part2(program)
}

func part1(program [][]string) {
	c := util.NewComputer()
	c.SetRegister("a", 7)
	c.Run(program)

	fmt.Printf("Computer state:\n%s\n", c)
}

func part2(program [][]string) {
	c := util.NewComputer()
	c.SetRegister("a", 12)
	c.Run(program)

	fmt.Printf("Computer state:\n%s\n", c)
}
