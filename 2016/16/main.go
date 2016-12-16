package main

import (
	"fmt"
)

func main() {
	disk := NewDisk(272)
	disk.Fill("01000100010010111")
	fmt.Printf("Checksum: %s\n", disk.Checksum())

	disk = NewDisk(35651584)
	disk.Fill("01000100010010111")
	fmt.Printf("Done with second fill...\n")
	fmt.Printf("Checksum: %s\n", disk.Checksum())
}
