package main

import (
	"fmt"
)

func main() {
	i, _ := FillString(272, "01000100010010111")
	c := Checksum(true, i)
	fmt.Printf("Checksum: %s\n", c)
	i, _ = FillString(35651584, "01000100010010111")
	fmt.Printf("Done with second fill...\n")
	fmt.Printf("Checksum: %s\n", c)
}
