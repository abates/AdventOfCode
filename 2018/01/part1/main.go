package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	var value int
	frequency := 0
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		fmt.Sscanf(line, "%d", &value)
		frequency += value
	}
	fmt.Printf("Final Frequency: %d\n", frequency)
}
