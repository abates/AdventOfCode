package main

import (
	"fmt"

	"github.com/abates/AdventOfCode/2017/hash"
)

func main() {
	lengths := []byte{14, 58, 0, 116, 179, 16, 1, 104, 2, 254, 167, 86, 255, 55, 122, 244}

	h := hash.Compute(1, lengths)
	fmt.Printf("Value: %d\n", int(h[0])*int(h[1]))
}
