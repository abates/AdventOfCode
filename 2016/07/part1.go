package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
)

func main() {
	count := 0
	for _, line := range util.ReadInput() {
		ip := ParseIP(line)
		if ip.SupportsTLS() {
			count++
		}

	}
	fmt.Printf("%d\n", count)
}
