package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func react(polymer []rune) []rune {
	for i := 1; i < len(polymer)-1; i++ {
		if i < 0 {
			continue
		}
		r := polymer[i]
		if unicode.IsLower(r) {
			r = unicode.ToUpper(r)
		} else {
			r = unicode.ToLower(r)
		}

		if r == polymer[i] {
			polymer = append(polymer[0:i-1], polymer[i+1:]...)
			i -= 2
		} else if r == polymer[i+1] {
			polymer = append(polymer[0:i], polymer[i+2:]...)
			i -= 2
		}
	}

	return polymer
}

func removeUnit(polymer []rune, unit rune) []rune {
	str := string(polymer)
	unit = unicode.ToLower(unit)
	str = strings.Replace(str, string([]rune{unit}), "", -1)
	unit = unicode.ToUpper(unit)
	return []rune(strings.Replace(str, string([]rune{unit}), "", -1))
}

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Error: " + err.Error())
	}

	polymer := []rune("")
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		polymer = append(polymer, []rune(line)...)
	}

	units := make(map[rune]struct{})
	for _, r := range polymer {
		units[unicode.ToLower(r)] = struct{}{}
	}

	min := -1
	minPolymer := []rune{}
	for r := range units {
		candidate := removeUnit(polymer, r)
		candidate = react(candidate)
		if min == -1 || len(candidate) < min {
			min = len(candidate)
			minPolymer = candidate
		}
	}

	fmt.Printf("Length: %d\n", len(minPolymer))
}
