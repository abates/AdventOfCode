package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abates/AdventOfCode/2018/tm"
)

func part1(input []byte) error {
	program := tm.Program{}
	err := program.UnmarshalText(input)
	if err == nil {
		machine := tm.NewMachine()
		machine.Execute(tm.Registers{0, 0, 0, 0, 0, 0}, program, nil)
		fmt.Printf("Part 1: %v\n", machine.Registers)
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

	err = part1(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Part 1 failed: %v\n", err)
		os.Exit(-1)
	}
}
