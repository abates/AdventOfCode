package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Indicator rune

func (i Indicator) String() string {
	return string([]rune{rune(i)})
}

type Pattern struct {
	indicators [5]Indicator
	nextGen    Indicator
}

func (p *Pattern) UnmarshalText(text []byte) error {
	runes := []rune(string(text))
	// no input validation for you!
	p.indicators[0] = Indicator(runes[0])
	p.indicators[1] = Indicator(runes[1])
	p.indicators[2] = Indicator(runes[2])
	p.indicators[3] = Indicator(runes[3])
	p.indicators[4] = Indicator(runes[4])
	p.nextGen = Indicator(runes[9])
	return nil
}

func (p *Pattern) String() string {
	return fmt.Sprintf("%v%v%v%v%v => %v", p.indicators[0], p.indicators[1], p.indicators[2], p.indicators[3], p.indicators[4], p.nextGen)
}

type Pot struct {
	num       int
	leftPot   *Pot
	rightPot  *Pot
	indicator Indicator
}

func (p *Pot) Left(n int) *Pot {
	return p.left(p, n)
}

func (p *Pot) left(previous *Pot, n int) *Pot {
	if p == nil {
		p = &Pot{num: previous.num - 1, indicator: '.'}
		p.rightPot = previous
	}

	if n == 0 {
		return p
	}
	return p.leftPot.left(p, n-1)
}

func (p *Pot) Right(n int) *Pot {
	return p.right(p, n)
}

func (p *Pot) right(previous *Pot, n int) *Pot {
	if p == nil {
		p = &Pot{num: previous.num + 1, indicator: '.'}
		p.leftPot = previous
	}

	if n == 0 {
		return p
	}
	return p.rightPot.right(p, n-1)
}

func (p *Pot) Set(patterns map[[5]Indicator]*Pattern, previous *Pot) {
	var key [5]Indicator
	key[0] = previous.Left(2).indicator
	key[1] = previous.Left(1).indicator
	key[2] = previous.indicator
	key[3] = previous.Right(1).indicator
	key[4] = previous.Right(2).indicator
	if pattern, found := patterns[key]; found {
		p.indicator = pattern.nextGen
	} else {
		fmt.Printf("WARNING: No pattern found for key %s\n", key)
		p.indicator = '.'
	}
	p.num = previous.num
}

type Row struct {
	firstPot *Pot
	lastPot  *Pot
}

func (row *Row) Append(pot *Pot) {
	if row.firstPot == nil {
		row.firstPot = pot
		row.lastPot = pot
	} else {
		row.lastPot.rightPot = pot
		pot.leftPot = row.lastPot
		row.lastPot = pot
	}
}

func (row *Row) UnmarshalText(text []byte) error {
	runes := []rune(string(text))
	for i, r := range runes {
		if r != '.' && r != '#' {
			return fmt.Errorf("Pot should either be '.' or '#' got %q", r)
		}
		row.Append(&Pot{num: i, indicator: Indicator(r)})
	}
	return nil
}

func (row *Row) Advance(patterns map[[5]Indicator]*Pattern) *Row {
	newRow := &Row{}

	for pot := row.firstPot.Left(2); pot != nil; pot = pot.rightPot {
		newPot := &Pot{num: pot.num}
		newPot.Set(patterns, pot)
		newRow.Append(newPot)
	}

	for i := 1; i <= 2; i++ {
		newPot := &Pot{}
		newPot.Set(patterns, row.lastPot.Right(i))
		newRow.Append(newPot)
	}

	for pot := newRow.firstPot; pot.indicator == '.'; pot = pot.rightPot {
		newRow.firstPot = pot
		pot.leftPot = nil
	}

	for pot := newRow.lastPot; pot.indicator == '.'; pot = pot.leftPot {
		newRow.lastPot = pot
		pot.rightPot = nil
	}

	return newRow
}

func (row *Row) String() string {
	runes := []rune{}
	for pot := row.firstPot; pot != nil; pot = pot.rightPot {
		runes = append(runes, rune(pot.indicator))
	}
	return string(runes)
}

func (row *Row) Sum(offset int) int {
	sum := 0

	for pot := row.firstPot; pot != nil; pot = pot.rightPot {
		if pot.indicator == '#' {
			sum += pot.num + offset
		}
	}

	return sum
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input file> <num generations>\n", os.Args[0])
		os.Exit(-1)
	}

	maxGenerations, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid number of generations: %v\n", err)
		os.Exit(-1)
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input file: %v\n", err)
		os.Exit(-1)
	}

	row := &Row{}
	patterns := make(map[[5]Indicator]*Pattern)

	for _, line := range bytes.Split(input, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		str := string(line)
		if strings.HasPrefix(str, "initial state: ") {
			str = str[len("initial state: "):]
			err = row.UnmarshalText([]byte(str))
			if err != nil {
				panic("Failed to unmarshal initial state: " + err.Error())
			}
		} else {
			pattern := &Pattern{}
			err = pattern.UnmarshalText(line)
			if err != nil {
				panic("Failed to unmarshal pattern: " + err.Error())
			}
			patterns[pattern.indicators] = pattern
		}
	}

	previous := row
	offset := 0
	delta := 0
	for i := 0; i < maxGenerations; i++ {
		row = row.Advance(patterns)
		if previous.String() == row.String() {
			offset = i
			delta = row.firstPot.num - previous.firstPot.num
			break
		}
		previous = row
	}
	sum := row.Sum((maxGenerations - offset - 1) * delta)
	fmt.Printf("Num Plants: %d\n", sum)
}
