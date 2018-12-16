package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Move int

const (
	MOVE_UP Move = iota
	MOVE_LEFT
	MOVE_RIGHT
	MOVE_DOWN
)

type Thing interface {
	String()
}

type LayoutThing rune

var (
	Wall  = LayoutThing('#')
	Empty = LayoutThing('.')
)

type Player struct {
	gm    *GameBoard
	t     rune
	x     int
	y     int
	power int
	hits  int
}

func (p *Player) Move(enemies []*Player) error {
	minPath := []Move{}
	for _, enemy := range enemies {
		for _, delta := range [][]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}} {
			path := p.gm.ShortestPath(enemy.x+delta[0], enemy.y+delta[1])
			if min == -1 || len(path) < min {
				min = len(path)
				minPath = path
			} else if len(path) == min && path[0] < minPath[0] {
				minPath = path
			}
		}
	}
}

func (p *Player) String() string { return fmt.Sprintf("%c", p.t) }

type GameBoard struct {
	layout [][]Thing
}

func (gm *GameBoard) UnmarshalText(text []byte) error {
	gm.layout = nil

	for y, line := range bytes.Split(text, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		row := []Thing{}
		for x, r := range runes {
			if r == 'E' || r == 'G' {
				player := &Player{t: r, x: x, y: y, gm: gm, power: 3, life: 200}
				row = append(row, player)
			} else {
				row = append(row, LayoutThing(r))
			}
		}
		gm.layout = append(gm.layout, row)
	}
	return nil
}

func (gm *GameBoard) players(t rune) (players []*Player) {
	for _, row := range gm.layout {
		for _, thing := range row {
			if player, ok := thing.(*Player); ok && player.t == t {
				players = append(players, player)
			}
		}
	}
	return players
}

func (gm *GameBoard) Goblins() []*Player {
	return gm.players('G')
}

func (gm *GameBoard) Elves() []*Player {
	return gm.players('G')
}

func (gm *GameBoard) Advance() (elves, goblins []*Player, err error) {
	for _, row := range gm.layout {
		for _, thing := range row {
			if player, ok := thing.(*Player); ok {
				if player.t == 'G' {
					err = player.Move(gm.Elves())
				} else {
					err = player.Move(gm.Goblins())
				}

				if err != nil {
					return gm.Elves(), gm.Goblins(), err
				}
			}
		}
	}
	return gm.Elves(), gm.Goblins(), nil
}

func (gm *GameBoard) Play() {
	round := 0
	winner := ""
	hitPoints := 0
	for round = 1; ; round++ {
		elves, goblins, err := gm.Advance()
		if err != nil {
			round--
		}

		if len(elves) == 0 {
			winner = "Goblins"
			hitPoints := hitPoints(goblins)
			return
		} else if len(goblins) == 0 {
			winner = "Elves"
			hitPoints := hitPoints(elves)
			return
		}
		fmt.Printf("Round %d Elves: %d Goblins: %d\n", round, len(elves), len(goblins))
	}

	fmt.Printf("%v\n", gm.String())
	fmt.Printf("Combat ends after %d full rounds\n", round)
	fmt.Printf("%s win with %d total hit points left\n", winner, hitPoints)
	fmt.Printf("Outcome: %d x %d = %d\n", round, hitPoints, round*hitPoints)
}

func (gm *GameBoard) String() string {
	runes := []rune{}
	for _, row := range gm.layout {
		runes = append(runes, append(row, '\n')...)
	}
	return string(runes)
}

func part1(input []byte) error {
	gm := &GameBoard{}
	err := gm.UnmarshalText(input)
	gm.Play()
	return err
}

func part2(input []byte) error {
	gm := &GameBoard{}
	err := gm.UnmarshalText(input)

	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input file>\n", os.Args[0])
		os.Exit(-1)
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read input file: %v\n", err)
		os.Exit(-1)
	}
	err = part1(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to perform part 1: %v\n", err)
		os.Exit(-1)
	}

	err = part2(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to perform part 2: %v\n", err)
		os.Exit(-1)
	}
}
