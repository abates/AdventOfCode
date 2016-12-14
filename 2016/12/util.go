package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

type Computer struct {
	registers map[string]int
	pc        int
}

func (c *Computer) Run(program [][]string) {
	c.pc = 0
	for c.pc < len(program) {
		operation := program[c.pc]
		switch operation[0] {
		case "cpy":
			if i, err := strconv.Atoi(operation[1]); err == nil {
				c.CpyInt(i, operation[2])
			} else {
				c.CpyValue(operation[1], operation[2])
			}
			c.pc++
		case "inc":
			c.Inc(operation[1])
			c.pc++
		case "dec":
			c.Dec(operation[1])
			c.pc++
		case "jnz":
			count, _ := strconv.Atoi(operation[2])
			if i, err := strconv.Atoi(operation[1]); err == nil {
				c.JnzInt(i, count)
			} else {
				c.JnzValue(operation[1], count)
			}
		}
	}
}

func (c *Computer) CpyValue(src, dst string) {
	c.CpyInt(c.registers[src], dst)
}

func (c *Computer) CpyInt(value int, dst string) {
	c.registers[dst] = value
}

func (c *Computer) Inc(register string) {
	c.registers[register] = c.registers[register] + 1
}

func (c *Computer) Dec(register string) {
	c.registers[register] = c.registers[register] - 1
}

func (c *Computer) JnzValue(register string, count int) {
	c.JnzInt(c.registers[register], count)
}

func (c *Computer) JnzInt(value int, count int) {
	if value != 0 {
		c.pc += count
	} else {
		c.pc++
	}
}

func (c *Computer) String() string {
	buffer := &bytes.Buffer{}
	registers := make([]string, 0)
	for r, _ := range c.registers {
		registers = append(registers, r)
	}

	sort.Strings(registers)

	for _, r := range registers {
		buffer.Write([]byte(fmt.Sprintf("%s: %d\n", r, c.registers[r])))
	}

	return buffer.String()
}
