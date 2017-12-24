package main

import "fmt"

func factor(num int) []int {
	factors := []int{}
	stop := 1 + num/2
	start := 1
	for start < stop {
		if num%start == 0 {
			factors = append(factors, start)
		}
		start++
	}
	return factors
}

func main() {
	h := 0
	for b := 108400; b <= 125400; b += 17 {
		factors := factor(b)
		if len(factors) > 1 {
			h++
		}
	}
	fmt.Printf("H: %d\n", h)
}
