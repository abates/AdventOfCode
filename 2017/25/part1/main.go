package main

import "fmt"

func main() {
	tape := NewTape()
	Run(tape, StateA, 12302209)
	fmt.Printf("Checksum: %d\n", tape.Checksum())
}
