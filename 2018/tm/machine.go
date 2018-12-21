package tm

import (
	"bytes"
	"fmt"
)

type Registers [6]int

func (r Registers) String() string {
	return fmt.Sprintf("[%d, %d, %d, %d, %d, %d]", r[0], r[1], r[2], r[3], r[4], r[5])
}

type Machine struct {
	IP        int
	ipr       int
	Registers Registers
	opIndex   map[string]func(a, b, c int)
}

func NewMachine() *Machine {
	computer := &Machine{}
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

func (computer *Machine) addr(a, b, c int) {
	computer.Registers[c] = computer.Registers[a] + computer.Registers[b]
}

func (computer *Machine) addi(a, b, c int) {
	computer.Registers[c] = computer.Registers[a] + b
}

func (computer *Machine) mulr(a, b, c int) {
	computer.Registers[c] = computer.Registers[a] * computer.Registers[b]
}

func (computer *Machine) muli(a, b, c int) {
	computer.Registers[c] = computer.Registers[a] * b
}

func (computer *Machine) banr(a, b, c int) {
	computer.Registers[c] = computer.Registers[a] & computer.Registers[b]
}

func (computer *Machine) bani(a, b, c int) {
	computer.Registers[c] = computer.Registers[a] & b
}

func (computer *Machine) borr(a, b, c int) {
	computer.Registers[c] = computer.Registers[a] | computer.Registers[b]
}

func (computer *Machine) bori(a, b, c int) {
	computer.Registers[c] = computer.Registers[a] | b
}

func (computer *Machine) setr(a, b, c int) {
	computer.Registers[c] = computer.Registers[a]
}

func (computer *Machine) seti(a, b, c int) {
	computer.Registers[c] = a
}

func (computer *Machine) gtir(a, b, c int) {
	if a > computer.Registers[b] {
		computer.Registers[c] = 1
	} else {
		computer.Registers[c] = 0
	}
}

func (computer *Machine) gtri(a, b, c int) {
	if computer.Registers[a] > b {
		computer.Registers[c] = 1
	} else {
		computer.Registers[c] = 0
	}
}

func (computer *Machine) gtrr(a, b, c int) {
	if computer.Registers[a] > computer.Registers[b] {
		computer.Registers[c] = 1
	} else {
		computer.Registers[c] = 0
	}
}

func (computer *Machine) eqir(a, b, c int) {
	if a == computer.Registers[b] {
		computer.Registers[c] = 1
	} else {
		computer.Registers[c] = 0
	}
}

func (computer *Machine) eqri(a, b, c int) {
	if computer.Registers[a] == b {
		computer.Registers[c] = 1
	} else {
		computer.Registers[c] = 0
	}
}

func (computer *Machine) eqrr(a, b, c int) {
	if computer.Registers[a] == computer.Registers[b] {
		computer.Registers[c] = 1
	} else {
		computer.Registers[c] = 0
	}
}

func (computer *Machine) execute(op string, a, b, c int) {
	computer.opIndex[op](a, b, c)
}

func (computer *Machine) Execute(registers Registers, program Program, trace func() bool) int {
	instructions := []Instruction{}
	// Compile the program
	for _, instruction := range program.instructions {
		newInstruction := instruction
		newInstruction.op = computer.opIndex[instruction.Op]
		instructions = append(instructions, newInstruction)
	}

	// Run the program
	computer.ipr = program.ipr
	computer.Registers = registers

	count := 0
	for 0 <= computer.IP && computer.IP < len(instructions) {
		computer.Registers[computer.ipr] = computer.IP
		instructions[computer.IP].execute()

		if trace != nil {
			if !trace() {
				break
			}
		}

		computer.IP = computer.Registers[computer.ipr]
		computer.IP++
		count++
	}
	return count
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
