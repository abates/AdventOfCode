package main

import (
	"github.com/abates/AdventOfCode/2016/util"
	"strings"
)

func main() {
	program := make([][]string, 0)
	for _, line := range util.ReadInput() {
		program = append(program, strings.Split(line, " "))
	}
	part1(program)
}

func part1(program [][]string) {
	for i := 0; ; i++ {
		print(i, " ")
		c := util.NewComputer()
		c.SetRegister("a", int64(i))
		go c.Run(program)
		last := int64(1)
		for v := range c.Stdout {
			if last == v {
				c.Stdin <- 0
				println("")
				break
			} else if last == 0 {
				last = 1
			} else {
				last = 0
			}
			print(v)
		}
	}
}
