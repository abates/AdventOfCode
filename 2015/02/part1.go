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

func computeArea(length, width, height int) int {
	return 2*length*width + 2*width*height + 2*height*length
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

func minSideArea(length, width, height int) int {
	return min(length*width, length*height, width*height)
}

func main() {
	f, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(f)
	area := 0
	for _, line := range strings.Split(string(b), "\n") {
		if line == "" {
			continue
		}
		length, width, height := parseLine(line)
		area += computeArea(length, width, height) + minSideArea(length, width, height)
	}
	fmt.Println(area)
}
