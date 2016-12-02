package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	keypad := NewKeypad(1, 1, [][]string{
		[]string{"0", "0", "1", "0", "0"},
		[]string{"0", "2", "3", "4", "0"},
		[]string{"5", "6", "7", "8", "9"},
		[]string{"0", "A", "B", "C", "0"},
		[]string{"0", "0", "D", "0", "0"},
	})

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
