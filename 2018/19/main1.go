package main

import (
	"fmt"
	"math"
)

func factor(number int) []int {
	factors := make(map[int]struct{})
	factors[1] = struct{}{}
	factors[number] = struct{}{}
	sqrt := int(math.Sqrt(float64(number)))
	for test := number - 1; test >= sqrt; test-- {
		if number%test == 0 {
			factors[test] = struct{}{}
			factors[number/test] = struct{}{}
		}
	}

	fa := []int{}
	for v := range factors {
		fa = append(fa, v)
	}
	return fa
}

func main() {
	sum := 0
	for _, v := range factor(1030) {
		sum += v
	}
	fmt.Printf("Part 1: %d\n", sum)

	sum = 0
	for _, v := range factor(10551430) {
		sum += v
	}
	fmt.Printf("Part 2: %d\n", sum)
}
