package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func containsN(str string, n int) bool {
	counts := make(map[rune]int)
	for _, c := range str {
		counts[c]++
	}

	for _, c := range counts {
		if c == n {
			return true
		}
	}
	return false
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	twos := 0
	threes := 0

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		if containsN(line, 3) {
			threes++
		}

		if containsN(line, 2) {
			twos++
		}
	}

	fmt.Printf("Twos: %d Threes: %d\n", twos, threes)
	fmt.Printf("Checksum: %d\n", twos*threes)
}
