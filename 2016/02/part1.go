package main

import (
	"fmt"
)

func main() {
	keypad := NewKeypad(1, 1, nil)
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
