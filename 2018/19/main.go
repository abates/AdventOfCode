package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type Registers [6]int

func (r Registers) String() string {
	return fmt.Sprintf("[%d, %d, %d, %d, %d, %d]", r[0], r[1], r[2], r[3], r[4], r[5])
}

type Computer struct {
	ip        int
	ipr       int
	registers Registers
	opIndex   map[string]func(a, b, c int)
}

func NewComputer() *Computer {
	computer := &Computer{}
	computer.opIndex = map[string]func(a, b, c int){
		"addr": computer.addr,
		"addi": computer.addi,
		"mulr": computer.mulr,
		"muli": computer.muli,
		"banr": computer.banr,
		"bani": computer.bani,
		"borr": computer.borr,
		"bori": computer.bori,
		"setr": computer.setr,
		"seti": computer.seti,
		"gtir": computer.gtir,
		"gtri": computer.gtri,
		"gtrr": computer.gtrr,
		"eqir": computer.eqir,
		"eqri": computer.eqri,
		"eqrr": computer.eqrr,
	}
	return computer
}

func (computer *Computer) addr(a, b, c int) {
	computer.registers[c] = computer.registers[a] + computer.registers[b]
}

func (computer *Computer) addi(a, b, c int) {
	computer.registers[c] = computer.registers[a] + b
}

func (computer *Computer) mulr(a, b, c int) {
	computer.registers[c] = computer.registers[a] * computer.registers[b]
}

func (computer *Computer) muli(a, b, c int) {
	computer.registers[c] = computer.registers[a] * b
}

func (computer *Computer) banr(a, b, c int) {
	computer.registers[c] = computer.registers[a] & computer.registers[b]
}

func (computer *Computer) bani(a, b, c int) {
	computer.registers[c] = computer.registers[a] & b
}

func (computer *Computer) borr(a, b, c int) {
	computer.registers[c] = computer.registers[a] | computer.registers[b]
}

func (computer *Computer) bori(a, b, c int) {
	computer.registers[c] = computer.registers[a] | b
}

func (computer *Computer) setr(a, b, c int) {
	computer.registers[c] = computer.registers[a]
}

func (computer *Computer) seti(a, b, c int) {
	computer.registers[c] = a
}

func (computer *Computer) gtir(a, b, c int) {
	if a > computer.registers[b] {
		computer.registers[c] = 1
	} else {
		computer.registers[c] = 0
	}
}

func (computer *Computer) gtri(a, b, c int) {
	if computer.registers[a] > b {
		computer.registers[c] = 1
	} else {
		computer.registers[c] = 0
	}
}

func (computer *Computer) gtrr(a, b, c int) {
	if computer.registers[a] > computer.registers[b] {
		computer.registers[c] = 1
	} else {
		computer.registers[c] = 0
	}
}

func (computer *Computer) eqir(a, b, c int) {
	if a == computer.registers[b] {
		computer.registers[c] = 1
	} else {
		computer.registers[c] = 0
	}
}

func (computer *Computer) eqri(a, b, c int) {
	if computer.registers[a] == b {
		computer.registers[c] = 1
	} else {
		computer.registers[c] = 0
	}
}

func (computer *Computer) eqrr(a, b, c int) {
	if computer.registers[a] == computer.registers[b] {
		computer.registers[c] = 1
	} else {
		computer.registers[c] = 0
	}
}

func (computer *Computer) execute(op string, a, b, c int) {
	computer.opIndex[op](a, b, c)
}

func (computer *Computer) Execute(registers Registers, program Program) {
	instructions := []Instruction{}
	// Compile the program
	for _, instruction := range program.instructions {
		newInstruction := instruction
		newInstruction.op = computer.opIndex[instruction.Op]
		instructions = append(instructions, newInstruction)
	}

	// Run the program
	computer.ipr = program.ipr
	computer.registers = registers

	for 0 <= computer.ip && computer.ip < len(instructions) {
		fmt.Printf("%02d %v %v ", computer.ip, computer.registers, instructions[computer.ip])
		computer.registers[computer.ipr] = computer.ip
		instructions[computer.ip].execute()
		computer.ip = computer.registers[computer.ipr]
		computer.ip++
		fmt.Printf("%v\n", computer.registers)
		//if computer.ip == 14 {
		//fmt.Printf("%v\n", computer.registers)
		//}
	}
}

type Instruction struct {
	Op string
	op func(a, b, c int)
	A  int
	B  int
	C  int
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s %d %d %d", i.Op, i.A, i.B, i.C)
}

func (i *Instruction) execute() {
	i.op(i.A, i.B, i.C)
}

func (i *Instruction) UnmarshalText(text []byte) error {
	_, err := fmt.Sscanf(string(text), "%s %d %d %d", &i.Op, &i.A, &i.B, &i.C)
	return err
}

type Program struct {
	ipr          int
	instructions []Instruction
}

func (p *Program) UnmarshalText(text []byte) error {
	lines := bytes.Split(text, []byte("\n"))
	_, err := fmt.Sscanf(string(lines[0]), "#ip %d", &p.ipr)
	if err == nil {
		for _, line := range lines[1:] {
			if len(line) == 0 {
				continue
			}
			instruction := Instruction{}
			err = instruction.UnmarshalText(line)
			if err == nil {
				p.instructions = append(p.instructions, instruction)
			} else {
				break
			}
		}
	}
	return err
}

func part1(input []byte) error {
	program := Program{}
	err := program.UnmarshalText(input)
	if err == nil {
		computer := NewComputer()
		computer.Execute(Registers{0, 0, 0, 0, 0, 0}, program)
		fmt.Printf("Part 1: %v\n", computer.registers)
	}
	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <input file>\n", os.Args[0])
		os.Exit(-1)
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input file: %v\n", err)
		os.Exit(-1)
	}

	for i, f := range []func([]byte) error{part1} {
		err := f(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Part %d failed: %v\n", i+1, err)
			os.Exit(-1)
		}
	}
}
