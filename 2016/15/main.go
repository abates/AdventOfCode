package main

import (
	"fmt"
)

func GetPachinko() *Pachinko {
	pachinko := NewPachinko()
	pachinko.AddWheel(11, 13)
	pachinko.AddWheel(0, 5)
	pachinko.AddWheel(11, 17)
	pachinko.AddWheel(0, 3)
	pachinko.AddWheel(2, 7)
	pachinko.AddWheel(17, 19)
	return pachinko
}

func main() {
	pachinko := GetPachinko()
	t := pachinko.Run()
	fmt.Printf("T: %d\n", t)

	pachinko = GetPachinko()
	pachinko.AddWheel(0, 11)

	t = pachinko.Run()
	fmt.Printf("T: %d\n", t)
}
