package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parse(str string) []int {
	values := make([]int, 0)
	for _, f := range strings.Fields(str) {
		value, _ := strconv.Atoi(f)
		values = append(values, value)
	}
	return values
}

type Computer struct {
	instructions []int
	pc           int
	clock        int
}

func NewComputer(instructions []int) *Computer {
	return &Computer{instructions: instructions}
}

func (c *Computer) Execute() {
	pc := 0
	for {
		c.clock++
		pc += c.instructions[c.pc]

		if c.instructions[c.pc] < 3 {
			c.instructions[c.pc]++
		} else {
			c.instructions[c.pc]--
		}

		if pc < 0 || len(c.instructions) <= pc {
			break
		}
		c.pc = pc
	}
}

func main() {
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	computer := NewComputer(parse(string(b)))
	computer.Execute()
	fmt.Printf("Steps: %d\n", computer.clock)
}
