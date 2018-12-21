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
		seen := make(map[int]struct{})
		lastSeen := 0
		machine.Execute(tm.Registers{0, 0, 0, 0, 0, 0}, program, func() bool {
			if machine.IP == 28 {
				if len(seen) == 0 {
					fmt.Printf("Part 1: %d\n", machine.Registers[4])
				} else if _, found := seen[machine.Registers[4]]; found {
					fmt.Printf("Part 2: %d\n", lastSeen)
					return false
				}
				lastSeen = machine.Registers[4]
				seen[machine.Registers[4]] = struct{}{}
			}
			return true
		})
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
