package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
)

type Int struct {
	*big.Int
}

func (i *Int) Addr() Address {
	return Address(i.Int64())
}

type Address int
type Opcode int
type Instruction [4]int
type Mode int

const (
	PositionMode  Mode = 0
	ImmediateMode      = 1
	RelativeMode       = 2
)

const (
	Add      Opcode = 1
	Multiply        = 2
	Input           = 3
	Output          = 4
	JIT             = 5
	JIF             = 6
	LT              = 7
	EQ              = 8
	ARB             = 9
	Return          = 99
)

func (inst *Instruction) String() string {
	return fmt.Sprintf("%1d%1d%1d%02d", inst[3], inst[2], inst[1], inst[0])
}

func (inst *Instruction) Opcode() Opcode {
	return Opcode(inst[0])
}

func (inst *Instruction) Mode(pos int) Mode {
	if pos < 1 {
		panic("position should be between 1 and 3 inclusive")
	}
	return Mode(inst[pos])
}

func (inst *Instruction) ParseInt(input *Int) (err error) {
	input = &Int{big.NewInt(0).Set(input.Int)}
	inst[0] = int(big.NewInt(0).Mod(input.Int, big.NewInt(100)).Int64())
	input.Div(input.Int, big.NewInt(100))
	for i := 1; i < 4 && input.Cmp(big.NewInt(0)) > 0; i++ {
		inst[i] = int(big.NewInt(0).Mod(input.Int, big.NewInt(10)).Int64())
		input.Div(input.Int, big.NewInt(10))
	}
	return nil
}

func ParseComputerMemory(lines []string) ([]*Int, error) {
	if len(lines) < 1 {
		return nil, fmt.Errorf("Expected one line in file")
	}

	mem := []*Int{}
	for _, digits := range strings.Split(lines[0], ",") {
		i := &Int{big.NewInt(0)}
		_, success := i.SetString(digits, 10)
		if !success {
			return nil, fmt.Errorf("Failed to parse %q", digits)
		}
		mem = append(mem, i)
	}
	return mem, nil
}

type Computer struct {
	name   string
	mem    []*Int
	output io.Writer
	input  io.Reader
	rb     Address
}

func NewComputer(mem []*Int) *Computer {
	c := &Computer{
		mem:    make([]*Int, len(mem)),
		input:  os.Stdin,
		output: os.Stdout,
	}

	copy(c.mem, mem)
	return c
}

func (c *Computer) SetInput(r io.Reader) {
	c.input = r
}

func (c *Computer) Input(mode Mode, addr Address) {
	v := 0
	fmt.Fscanf(c.input, "%d", &v)
	c.Set(mode, addr, &Int{big.NewInt(int64(v))})
}

func (c *Computer) SetOutput(w io.Writer) {
	c.output = w
}

func (c *Computer) Output(i *Int) {
	fmt.Fprintf(c.output, "%s\n", i.String())
}

func (c *Computer) Get(mode Mode, addr Address) *Int {
	if mode == PositionMode {
		addr = c.mem[addr].Addr()
	} else if mode == RelativeMode {
		addr = c.rb + c.mem[addr].Addr()
	}

	if int(addr) >= len(c.mem) {
		return &Int{big.NewInt(0)}
	}
	return c.mem[addr]
}

func (c *Computer) Set(mode Mode, addr Address, value *Int) {
	if mode == PositionMode {
		addr = c.mem[addr].Addr()
	} else if mode == RelativeMode {
		addr = c.rb + c.mem[addr].Addr()
	}

	if int(addr) >= len(c.mem) {
		c.mem = append(c.mem, make([]*Int, int(addr)-len(c.mem)+1)...)
	}
	c.mem[addr] = value
}

func (c *Computer) Add(x, y *Int) *Int {
	v := &Int{big.NewInt(0)}
	v.Add(x.Int, y.Int)
	return v
}

func (c *Computer) Multiply(x, y *Int) *Int {
	v := &Int{big.NewInt(0)}
	v.Mul(x.Int, y.Int)
	return v
}

func (c *Computer) RunWithInput(input string) (output string, err error) {
	r := strings.NewReader(fmt.Sprintf("%s\n", input))
	w := &strings.Builder{}
	c.SetInput(r)
	c.SetOutput(w)
	err = c.Run()
	return strings.TrimSpace(w.String()), err
}

func (c *Computer) Run() error {
	for i := 0; i < len(c.mem); {
		inst := Instruction{}
		inst.ParseInt(c.mem[i])
		if inst.Opcode() == Return {
			return nil
		}

		//println(i, c.Dump())
		switch inst.Opcode() {
		case Add:
			c.Set(inst.Mode(3), Address(i+3), c.Add(c.Get(inst.Mode(1), Address(i+1)), c.Get(inst.Mode(2), Address(i+2))))
			i += 4
		case Multiply:
			c.Set(inst.Mode(3), Address(i+3), c.Multiply(c.Get(inst.Mode(1), Address(i+1)), c.Get(inst.Mode(2), Address(i+2))))
			i += 4
		case Input:
			c.Input(inst.Mode(1), Address(i+1))
			i += 2
		case Output:
			c.Output(c.Get(inst.Mode(1), Address(i+1)))
			i += 2
		case JIT:
			if c.Get(inst.Mode(1), Address(i+1)).Cmp(big.NewInt(0)) == 0 {
				i += 3
			} else {
				i = int(c.Get(inst.Mode(2), Address(i+2)).Int64())
			}
		case JIF:
			if c.Get(inst.Mode(1), Address(i+1)).Cmp(big.NewInt(0)) == 0 {
				i = int(c.Get(inst.Mode(2), Address(i+2)).Int64())
			} else {
				i += 3
			}
		case LT:
			if c.Get(inst.Mode(1), Address(i+1)).Cmp(c.Get(inst.Mode(2), Address(i+2)).Int) < 0 {
				c.Set(inst.Mode(3), Address(i+3), &Int{big.NewInt(1)})
			} else {
				c.Set(inst.Mode(3), Address(i+3), &Int{big.NewInt(0)})
			}
			i += 4
		case EQ:
			if c.Get(inst.Mode(1), Address(i+1)).Cmp(c.Get(inst.Mode(2), Address(i+2)).Int) == 0 {
				c.Set(inst.Mode(3), Address(i+3), &Int{big.NewInt(1)})
			} else {
				c.Set(inst.Mode(3), Address(i+3), &Int{big.NewInt(0)})
			}
			i += 4
		case ARB:
			c.rb += c.Get(inst.Mode(1), Address(i+1)).Addr()
			i += 2
		default:
			return fmt.Errorf("Encountered unknown opcode %d", inst.Opcode())
		}
	}
	return nil
}

func (c *Computer) Dump() string {
	str := []string{}
	for _, v := range c.mem {
		str = append(str, fmt.Sprintf("%d", v))
	}
	return strings.Join(str, ",")
}
