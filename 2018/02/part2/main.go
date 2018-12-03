package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func diff(str1, str2 string) int {
	count := 0
	runes := []rune(str2)
	for i, c := range str1 {
		if c != runes[i] {
			count++
		}
	}
	return count
}

func slice(str1, str2 string) string {
	output := []rune{}
	runes := []rune(str2)
	for i, c := range str1 {
		if c == runes[i] {
			output = append(output, c)
		}
	}
	return string(output)
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")

	ids := []string{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		ids = append(ids, line)
	}

	for i, id1 := range ids {
		for _, id2 := range ids[i+1:] {
			if diff(id1, id2) == 1 {
				fmt.Printf("IDs: %s %s\n", id1, id2)
				fmt.Printf("Slice: %s\n", slice(id1, id2))
			}
		}
	}
}
