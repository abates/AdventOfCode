package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

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

	for i := 1; i < len(polymer)-1; i++ {
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

	fmt.Printf("Length: %d\n", len(polymer))
}
