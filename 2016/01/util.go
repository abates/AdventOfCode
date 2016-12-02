package main

import (
	"bufio"
	"strconv"
	"strings"
)

func readInput(reader *bufio.Reader, cb func(string, int) bool) error {
	line, _, err := reader.ReadLine()
	if len(line) == 0 || err != nil {
		return err
	}

	for _, operation := range strings.Split(string(line), ", ") {
		tokens := strings.SplitN(operation, "", 2)
		distance, _ := strconv.Atoi(tokens[1])
		if !cb(tokens[0], distance) {
			break
		}
	}
	return nil
}
