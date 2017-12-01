package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func parseLine(line string) (length, width, height int) {
	tokens := strings.Split(line, "x")
	return atoi(tokens[0]), atoi(tokens[1]), atoi(tokens[2])
}

func volume(length, width, height int) int {
	return length * width * height
}

func min(i ...int) int {
	a := i[0]
	b := 0
	if len(i) == 2 {
		b = i[1]
	} else {
		b = min(i[1:]...)
	}

	if a < b {
		return a
	}
	return b
}

func perimeter(length, width int) int {
	return 2*length + 2*width
}

func minPerim(length, width, height int) int {
	return min(perimeter(length, width), perimeter(width, height), perimeter(length, height))
}

func main() {
	f, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(f)
	total := 0
	for _, line := range strings.Split(string(b), "\n") {
		if line == "" {
			continue
		}
		length, width, height := parseLine(line)

		total += minPerim(length, width, height) + volume(length, width, height)
	}
	fmt.Println(total)
}
