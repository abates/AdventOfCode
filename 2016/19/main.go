package main

import (
	"fmt"
)

func main() {
	e := NewWhiteElephantExchange(3017957, TakeGiftsLeft)
	id := e.Exchange()

	fmt.Printf("Elf ID: %d\n", id)

	e = NewWhiteElephantExchange(3017957, TakeGiftsAcross)
	id = e.Exchange()

	fmt.Printf("Elf ID: %d\n", id)
}
