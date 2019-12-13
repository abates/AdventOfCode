package main

import (
	"strings"
)

func init() {
	d5 := &D5{}
	challenges = append(challenges, &challenge{"Day 05", "input/day05.txt", nil, d5.parseFile, d5.part1, d5.part2})
}

type D5 struct {
	mem []*Int
}

func (d5 *D5) parseFile(lines []string) (err error) {
	d5.mem, err = ParseComputerMemory(lines)
	return err
}

func (d5 *D5) run(input string) (string, error) {
	output, err := NewComputer(d5.mem).RunWithInput(input)
	o := strings.Split(output, "\n")
	return o[len(o)-1], err
}

func (d5 *D5) part1() (string, error) {
	return d5.run("1")
}

func (d5 *D5) part2() (string, error) {
	return d5.run("5")
}
