package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/abates/AdventOfCode/util"
)

func main() {
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	sum := 0
	for _, line := range strings.Split(string(b), "\n") {
		values := make([]int, 0)
		for _, field := range strings.Fields(line) {
			value, _ := strconv.Atoi(field)
			values = append(values, value)
		}
		max := util.Max(values...)
		min := util.Min(values...)
		sum += max - min
	}
	fmt.Printf("Checksum: %d\n", sum)
}
