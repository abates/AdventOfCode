package main

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	d4 := &D4{}
	challenges[4] = &challenge{"Day 04", "input/day04.txt", d4}
}

type D4 struct {
	min int
	max int
}

func (d4 *D4) parseFile(lines []string) (err error) {
	if len(lines) < 1 {
		return fmt.Errorf("Input file should have only one line.  File had %d lines", len(lines))
	}
	r := strings.Split(lines[0], "-")
	if len(r) != 2 {
		return fmt.Errorf("Expected input format to be <integer>-<integer> %q does not match", lines[0])
	}

	d4.min, err = strconv.Atoi(r[0])
	if err == nil {
		d4.max, err = strconv.Atoi(r[1])
	}
	return err
}

func splitInt(places, input int) []int {
	parts := make([]int, places)
	for i := 0; input > 0; i++ {
		parts[i] = input % 10
		input = input / 10
	}

	// reverse the array
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return parts
}

func checkIncrementing(parts []int) bool {
	last := parts[0]
	for _, part := range parts[1:] {
		if last > part {
			return false
		}
		last = part
	}
	return true
}

func checkRepeating(parts []int, doubleOnly bool) bool {
	found := make([]int, 10)
	last1 := parts[0]

	for _, part := range parts[1:] {
		if last1 == part {
			found[part]++
		}
		last1 = part
	}

	for _, f := range found {
		if f >= 1 {
			if doubleOnly {
				if f == 1 {
					return true
				}
			} else {
				return true
			}
		}
	}
	return false
}

func (d4 *D4) part1() (string, error) {
	count := 0
	for i := d4.min; i <= d4.max; i++ {
		parts := splitInt(6, i)
		if checkIncrementing(parts) && checkRepeating(parts, false) {
			count++
		}
	}
	return fmt.Sprintf("There are %d passwords that match the criteria", count), nil
}

func (d4 *D4) part2() (string, error) {
	count := 0
	for i := d4.min; i <= d4.max; i++ {
		parts := splitInt(6, i)
		if checkIncrementing(parts) {
			if checkRepeating(parts, true) {
				count++
			}
		}
	}
	return fmt.Sprintf("There are %d passwords that match the criteria", count), nil
}
