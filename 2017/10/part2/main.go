package main

import (
	"fmt"

	"github.com/abates/AdventOfCode/2017/hash"
)

func main() {
	h := hash.ComputeString("14,58,0,116,179,16,1,104,2,254,167,86,255,55,122,244")
	fmt.Printf("Hash: %q\n", h.String())
}
