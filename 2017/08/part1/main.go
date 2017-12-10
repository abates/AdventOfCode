package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Computer struct {
	max       int
	registers map[string]int
}

func (c *Computer) assign(register string, value int) {
	c.registers[register] = value
	if c.max < c.registers[register] {
		c.max = c.registers[register]
	}
}

func (c *Computer) Inc(register string, value int) {
	c.assign(register, c.registers[register]+value)
}

func (c *Computer) Dec(register string, value int) {
	c.assign(register, c.registers[register]-value)
}

func (c *Computer) Execute(statement string) {
	var register1 string
	var value1 int
	var operation string
	var register2 string
	var operator string
	var value2 int

	fmt.Sscanf(statement, "%s %s %d if %s %s %d", &register1, &operation, &value1, &register2, &operator, &value2)
	condition := false
	switch operator {
	case "<":
		condition = c.registers[register2] < value2
	case "<=":
		condition = c.registers[register2] <= value2
	case ">":
		condition = c.registers[register2] > value2
	case ">=":
		condition = c.registers[register2] >= value2
	case "!=":
		condition = c.registers[register2] != value2
	case "==":
		condition = c.registers[register2] == value2
	default:
		panic("Unknown operator " + operator + " in " + statement)
	}

	if condition {
		if operation == "inc" {
			c.Inc(register1, value1)
		} else if operation == "dec" {
			c.Dec(register1, value1)
		} else {
			panic("Unknown operation " + operation)
		}
	}
}

func main() {
	computer := &Computer{
		registers: make(map[string]int),
	}
	b, _ := ioutil.ReadAll(os.Stdin)
	for _, statement := range strings.Split(string(b), "\n") {
		if strings.TrimSpace(statement) == "" {
			continue
		}
		computer.Execute(statement)
	}

	found := false
	max := 0

	for _, value := range computer.registers {
		if !found || max < value {
			max = value
		}
		found = true
	}
	fmt.Printf("MAX is: %d\n", max)
	fmt.Printf("Runtime MAX is: %d\n", computer.max)
}
