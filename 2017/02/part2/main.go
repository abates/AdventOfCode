package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

		for i, v1 := range values {
			for _, v2 := range values[i+1:] {
				if v1%v2 == 0 {
					sum += v1 / v2
				} else if v2%v1 == 0 {
					sum += v2 / v1
				}
			}
		}
	}
	fmt.Printf("Checksum: %d\n", sum)
}
