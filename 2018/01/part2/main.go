package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	var value int
	adjustments := []int{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		fmt.Sscanf(line, "%d", &value)
		adjustments = append(adjustments, value)
	}

	frequencies := make(map[int]bool)
	curFreq := 0
	i := 0
	for {
		curFreq += adjustments[i]
		if _, found := frequencies[curFreq]; found {
			break
		}
		frequencies[curFreq] = true

		i++
		if i == len(adjustments) {
			i = 0
		}
	}
	fmt.Printf("Final Frequency: %d\n", curFreq)
}
