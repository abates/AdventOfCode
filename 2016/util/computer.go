package util

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Computer struct {
	registers map[string]int64
	program   [][]string
	pc        int
}

func NewComputer() *Computer {
	return &Computer{
		registers: map[string]int64{
			"a": 0,
			"b": 0,
			"c": 0,
			"d": 0,
		},
	}
}

func (c *Computer) value(v string) int64 {
	if i, err := strconv.Atoi(v); err == nil {
		return int64(i)
	}
	return c.registers[v]
}

func (c *Computer) Run(program [][]string) {
	c.program = program
	c.pc = 0
	for c.pc < len(c.program) {
		operation := c.program[c.pc]
		//fmt.Printf("%d %s %v\n", c.pc, strings.Join(operation, " "), c.registers)
		switch operation[0] {
		case "cpy":
			c.Cpy(c.value(operation[1]), operation[2])
			c.pc++
		case "inc":
			c.Inc(operation[1])
			c.pc++
		case "dec":
			c.Dec(operation[1])
			c.pc++
		case "jnz":
			c.Jnz(c.value(operation[1]), c.value(operation[2]))
		case "tgl":
			c.Tgl(c.value(operation[1]))
			c.pc++
		case "mul":
			c.Mul(c.value(operation[1]), c.value(operation[2]), operation[3])
			c.pc++
		case "noop":
			c.pc++
		default:
			panic(fmt.Sprintf("Unknown instruction %s", strings.Join(operation, " ")))
		}
	}
}

func (c *Computer) Mul(x, y int64, dst string) {
	c.SetRegister(dst, x*y)
}

func (c *Computer) Tgl(offset int64) {
	x := int64(c.pc) + offset
	if x >= int64(len(c.program)) {
		return
	}

	switch len(c.program[x]) {
	case 2:
		if c.program[x][0] == "inc" {
			c.program[x][0] = "dec"
		} else {
			c.program[x][0] = "inc"
		}
	case 3:
		if c.program[x][0] == "jnz" {
			c.program[x][0] = "cpy"
		} else {
			c.program[x][0] = "jnz"
		}
	default:
		panic("Don't know how to handle more than two arguments")
	}
}

func (c *Computer) SetRegister(register string, value int64) {
	if _, found := c.registers[register]; found {
		c.registers[register] = value
	}
}

func (c *Computer) Cpy(src int64, dst string) {
	c.SetRegister(dst, src)
}

func (c *Computer) Inc(register string) {
	c.SetRegister(register, c.registers[register]+1)
}

func (c *Computer) Dec(register string) {
	c.SetRegister(register, c.registers[register]-1)
}

func (c *Computer) Jnz(value int64, count int64) {
	//fmt.Printf("Jumping %d from %d to %d\n", count, c.pc, c.pc+count)
	if value != 0 {
		c.pc += int(count)
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
