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
	for _, c := range string(b) {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}
	fmt.Println(floor)
}
