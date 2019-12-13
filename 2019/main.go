package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type challenge struct {
	name       string
	inputFile  string
	lineParser func(string) error
	fileParser func([]string) error
	part1      func() (string, error)
	part2      func() (string, error)
}

var challenges []*challenge

func runChallenge(w io.Writer, r io.Reader, challenge *challenge, part int) error {
	input, err := ioutil.ReadAll(r)
	if err == nil {
		lines := strings.Split(string(input), "\n")
		if challenge.lineParser != nil {
			for _, line := range lines {
				if line == "" {
					continue
				}
				err = challenge.lineParser(line)
				if err != nil {
					break
				}
			}
		} else if challenge.fileParser != nil {
			err = challenge.fileParser(lines)
		} else {
			return fmt.Errorf("No parsing mechanism provided for challenge")
		}

		if err == nil {
			answer := ""
			if part&0x01 == 0x01 {
				answer, err = challenge.part1()
				fmt.Fprintf(w, "%s Part 1: %s\n", challenge.name, answer)
			}

			if err == nil {
				if part&0x02 == 2 {
					answer, err = challenge.part2()
					if err == nil {
						fmt.Fprintf(w, "%s Part 2: %s\n", challenge.name, answer)
					}
				}
			}
		}
	}
	return err
}

func usage(w io.Writer, msg string) {
	if msg != "" {
		fmt.Fprintf(w, "%s\n", msg)
	}
	fmt.Fprintf(w, "Usage: %v [challenge number]\n", os.Args[0])
	fmt.Fprintf(w, "  omit challenge number to run all challenges\n")
	os.Exit(-1)
}

func execute(c int) error {
	if c < 0 || len(challenges) < c {
		usage(os.Stderr, fmt.Sprintf("Challenge index out of bounds, must be between 0 and %d (inclusive)", len(challenges)-1))
	}

	challenge := challenges[c]
	file, err := os.Open(challenge.inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s - Failed opening input file: %v\n", challenge.name, err)
		os.Exit(-1)
	}

	return runChallenge(os.Stdout, file, challenge, 3)
}

func main() {
	var err error
	if len(os.Args) > 1 {
		c, err := strconv.Atoi(os.Args[1])
		if err == nil {
			err = execute(c)
			os.Exit(0)
		} else {
			usage(os.Stderr, fmt.Sprintf("Failed to parse challenge number: %v", err))
		}
	}

	for c := range challenges {
		err = execute(c)
		if err != nil {
			break
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
}
