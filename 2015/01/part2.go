package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(file)
	floor := 0
	for i, c := range string(b) {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}

		if floor == -1 {
			fmt.Println(i + 1)
			break
		}
	}
}
