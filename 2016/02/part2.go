package main

import (
	"fmt"
)

func main() {
	keypad := NewKeypad(0, 2, [][]string{
		[]string{"0", "0", "1", "0", "0"},
		[]string{"0", "2", "3", "4", "0"},
		[]string{"5", "6", "7", "8", "9"},
		[]string{"0", "A", "B", "C", "0"},
		[]string{"0", "0", "D", "0", "0"},
	})

	for _, sequence := range readSequences() {
		for _, direction := range sequence {
			err := keypad.Move(direction)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				break
			}
		}
		fmt.Printf("%s", keypad.Position())
	}
	fmt.Println("")
}
