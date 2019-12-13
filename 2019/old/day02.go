package main

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	d2 := &D2{}
	challenges = append(challenges, &challenge{"Day 2", "input/day2.txt", nil, d2.parse, d2.part1, d2.part2})
}

type D2 struct {
	mem []int
}

func (d2 *D2) parse(lines []string) error {
	if len(lines) < 1 {
		return fmt.Errorf("Expected one line in file")
	}

	for _, digits := range strings.Split(lines[0], ",") {
		v, err := strconv.Atoi(digits)
		if err != nil {
			return err
		}
		d2.mem = append(d2.mem, v)
	}
	return nil
}

func (d2 *D2) runProgram(mem []int) error {
	for i := 0; i < len(mem); i += 4 {
		op := mem[i]
		if op != 1 && op != 2 && op != 99 {
			return fmt.Errorf("Encountered unknown opcode %d", op)
		}

		if op == 99 {
			return nil
		}

		a := mem[mem[i+1]]
		b := mem[mem[i+2]]
		c := mem[i+3]

		if op == 1 {
			mem[c] = a + b
		} else {
			mem[c] = a * b
		}
	}
	return nil
}

func (d2 *D2) dump(mem []int) string {
	str := []string{}
	for _, v := range mem {
		str = append(str, fmt.Sprintf("%d", v))
	}
	return strings.Join(str, ",")
}

func (d2 *D2) part1() (string, error) {
	mem := make([]int, len(d2.mem))
	copy(mem, d2.mem)

	mem[1] = 12
	mem[2] = 2
	err := d2.runProgram(mem)
	answer := ""
	if err == nil {
		answer = fmt.Sprintf("Computer memory position 0 is %d", mem[0])
	}
	return answer, err
}

func (d2 *D2) part2() (string, error) {
	mem := make([]int, len(d2.mem))
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			copy(mem, d2.mem)

			mem[1] = a
			mem[2] = b
			err := d2.runProgram(mem)
			if err != nil {
				return "", err
			}

			if mem[0] == 19690720 {
				return fmt.Sprintf("noun=%d verb=%d: %d", a, b, 100*a+b), nil
			}
		}
	}
	return "", fmt.Errorf("No solution found")
}
