package main

import (
	"fmt"
	"math/big"
)

func init() {
	d2 := &D2{}
	challenges[2] = &challenge{"Day 02", "input/day02.txt", d2}
}

type D2 struct {
	mem []*Int
}

func (d2 *D2) parseFile(lines []string) (err error) {
	d2.mem, err = ParseComputerMemory(lines)
	return err
}

func (d2 *D2) part1() (string, error) {
	c := NewComputer(d2.mem)
	c.Set(ImmediateMode, 1, &Int{big.NewInt(12)})
	c.Set(ImmediateMode, 2, &Int{big.NewInt(2)})
	err := c.Run()
	answer := ""
	if err == nil {
		answer = fmt.Sprintf("Computer memory position 0 is %d", c.Get(ImmediateMode, 0))
	}
	return answer, err
}

func (d2 *D2) part2() (string, error) {
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			c := NewComputer(d2.mem)
			c.Set(ImmediateMode, 1, &Int{big.NewInt(int64(a))})
			c.Set(ImmediateMode, 2, &Int{big.NewInt(int64(b))})
			err := c.Run()
			if err != nil {
				return "", err
			}

			if c.Get(ImmediateMode, 0).Cmp(big.NewInt(19690720)) == 0 {
				return fmt.Sprintf("noun=%d verb=%d: %d", a, b, 100*a+b), nil
			}
		}
	}
	return "", fmt.Errorf("Failed to find solution")
}
