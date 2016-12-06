package util

import (
	"bufio"
	"os"
)

func ReadInput() []string {
	lines := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		// skip blank lines
		if len(line) == 0 {
			continue
		}

		lines = append(lines, string(line))
	}
	return lines
}
