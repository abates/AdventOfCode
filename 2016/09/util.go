package main

import (
	"fmt"
	"strings"
)

type V1EncodedString string

func (s V1EncodedString) Decode() int {
	numBytes := 0
	input := string(s)
	start := 0
	pos := 0
	for ; pos < len(input); pos++ {
		if input[pos] == '(' {
			if start < pos {
				numBytes += pos - start
			}
			var length int
			var repeat int
			index := strings.Index(input[pos:], ")")
			if index == -1 {
				panic("No closing parenthesis!")
			}
			n, _ := fmt.Sscanf(input[pos:pos+index+1], "(%dx%d)", &length, &repeat)
			if n == 2 {
				numBytes += repeat * ((pos + index + length + 1) - (pos + index + 1))
				pos = pos + index + length
				start = pos + 1
			} else {
				panic(fmt.Sprintf("Invalid input at %s %d", input[pos:], n))
			}
		}
	}

	if start < len(input) {
		numBytes += pos - start
	}

	return numBytes
}

type V2EncodedString string

func V2Decode(input string) int {
	numBytes := 0
	start := 0
	pos := 0
	for ; pos < len(input); pos++ {
		if input[pos] == '(' {
			if start < pos {
				numBytes += pos - start
			}
			var length int
			var repeat int
			n, _ := fmt.Sscanf(input[pos:], "(%dx%d)", &length, &repeat)
			if n == 2 {
				index := strings.Index(input[pos:], ")")
				tokenLen := V2Decode(input[pos+index+1 : pos+index+length+1])
				numBytes += tokenLen * repeat
				pos = pos + index + length
				start = pos + 1
			} else {
				panic(fmt.Sprintf("Invalid input at %s %d", input[pos:], n))
			}
		}
	}

	if start < len(input) {
		numBytes += pos - start
	}
	return numBytes
}

func (s V2EncodedString) Decode() int {
	return V2Decode(string(s))
}
