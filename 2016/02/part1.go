package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	keypad := NewKeypad(1, 1, nil)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		if len(line) == 0 {
			continue
		}

		sequence := strings.Split(string(line), "")
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
