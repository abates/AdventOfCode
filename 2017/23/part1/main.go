package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Processor struct {
	registers map[string]int
	numMul    int
}

func (p *Processor) Reset() {
	p.registers = make(map[string]int)
	p.numMul = 0
}

func (p *Processor) getValue(input string) int {
	if value, err := strconv.Atoi(input); err == nil {
		return value
	}
	return p.registers[input]
}

func (p *Processor) Set(x string, y int) {
	p.registers[x] = y
}

func (p *Processor) Get(x string) int {
	return p.registers[x]
}

func (p *Processor) Sub(x string, y int) {
	p.registers[x] -= y
	if x == "h" {
		fmt.Printf("H: %d\n", p.registers[x])
	}
}

func (p *Processor) Mul(x string, y int) {
	p.registers[x] *= y
	p.numMul++
}

func (p *Processor) Run(input string) {
	program := []string{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		program = append(program, strings.TrimSpace(line))
	}

	pc := 0
	for 0 <= pc && pc < len(program) {
		fields := strings.Fields(program[pc])
		switch fields[0] {
		case "set":
			p.Set(fields[1], p.getValue(fields[2]))
		case "sub":
			p.Sub(fields[1], p.getValue(fields[2]))
		case "mul":
			p.Mul(fields[1], p.getValue(fields[2]))
		case "jnz":
			if p.getValue(fields[1]) != 0 {
				//if pc == 15 {
				//fmt.Printf("15: B: %d C: %d D: %d E: %d G: %d\n", p.Get("b"), p.Get("c"), p.Get("d"), p.Get("e"), p.Get("g"))
				//if pc == 25 {
				//fmt.Printf("25: B: %d C: %d D: %d E: %d\n", p.Get("b"), p.Get("c"), p.Get("d"), p.Get("e"))
				if pc == 23 {
					fmt.Printf("23: B: %d C: %d D: %d E: %d\n", p.Get("b"), p.Get("c"), p.Get("d"), p.Get("e"))
				}
				pc += p.getValue(fields[2])
				continue
			}
		}
		pc++
	}
}

func main() {
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)

	p := &Processor{registers: make(map[string]int)}
	//p.Run(string(b))
	//fmt.Printf("Part 1: %d\n", p.numMul)

	p.Reset()

	p.Set("a", 1)
	p.Set("b", 5000)
	p.Run(string(b))
	fmt.Printf("Part 2: %d\n", p.Get("h"))
}
