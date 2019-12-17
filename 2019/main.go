package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type lineParser interface {
	parse(string) error
}

type fileParser interface {
	parseFile([]string) error
}

type implementation interface {
	part1() (string, error)
	part2() (string, error)
}

type challenge struct {
	name      string
	inputFile string
	challenge implementation
}

var challenges = make(map[int]*challenge)

func parseFile(content string, challenge *challenge) (err error) {
	lines := strings.Split(content, "\n")
	if f, ok := challenge.challenge.(lineParser); ok {
		for _, line := range lines {
			if line == "" {
				continue
			}
			err = f.parse(line)
			if err != nil {
				break
			}
		}
	} else if f, ok := challenge.challenge.(fileParser); ok {
		err = f.parseFile(lines)
	} else {
		return fmt.Errorf("No parsing mechanism provided for challenge")
	}
	return err
}

func runChallenge(w io.Writer, r io.Reader, challenge *challenge, part int) error {
	input, err := ioutil.ReadAll(r)
	if err == nil {
		err = parseFile(string(input), challenge)
	}

	if err == nil {
		answer := ""
		if part&0x01 == 0x01 {
			answer, err = challenge.challenge.part1()
			fmt.Fprintf(w, "%s Part 1: %s\n", challenge.name, answer)
		}

		if err == nil {
			if part&0x02 == 2 {
				answer, err = challenge.challenge.part2()
				if err == nil {
					fmt.Fprintf(w, "%s Part 2: %s\n", challenge.name, answer)
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
	challenge, found := challenges[c]
	if !found {
		usage(os.Stderr, fmt.Sprintf("No challenge found for day %d", c))
	}

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

	days := []int{}
	for k := range challenges {
		days = append(days, k)
	}
	sort.Ints(days)
	for _, day := range days {
		err = execute(day)
		if err != nil {
			break
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
}
