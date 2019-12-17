package main

import (
	"fmt"
	"strconv"
)

func init() {
	d1 := &D1{}
	challenges[1] = &challenge{"Day 01", "input/day01.txt", d1}
}

// MFR is the Module Fuel Requirement function
func MFR(mass int) int {
	return (mass / 3) - 2
}

type D1 struct {
	modules []int
}

func (d1 *D1) parse(line string) error {
	v, err := strconv.Atoi(line)
	if err == nil {
		d1.modules = append(d1.modules, v)
	}
	return err
}

func (d1 *D1) part1() (string, error) {
	sum := 0
	for _, module := range d1.modules {
		sum += MFR(module)
	}
	return fmt.Sprintf("Sum of fuel requirements is %d", sum), nil
}

func TotalMFR(mass int) int {
	req := MFR(mass)
	if req > 0 {
		req += TotalMFR(req)
	} else {
		req = 0
	}
	return req
}

func (d1 *D1) part2() (string, error) {
	sum := 0
	for _, module := range d1.modules {
		sum += TotalMFR(module)
	}
	return fmt.Sprintf("Sum of fuel requirements is %d", sum), nil
}
