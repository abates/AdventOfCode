package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/bfs"
)

/*
 * The first floor contains a thulium generator, a thulium-compatible microchip, a plutonium generator, and a strontium generator.
 * The second floor contains a plutonium-compatible microchip and a strontium-compatible microchip.
 * The third floor contains a promethium generator, a promethium-compatible microchip, a ruthenium generator, and a ruthenium-compatible microchip.
 * The fourth floor contains nothing relevant.
 */
func part1() {
	initialState := &State{
		elevator: 0,
		levels: [][]string{
			{"TG", "TM", "PG", ".", "SG", ".", ".", ".", ".", "."},
			{".", ".", ".", "PM", ".", "SM", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "KG", "KM", "RG", "RM"},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		},
	}

	endState := &State{
		elevator: 3,
		levels: [][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{"TG", "TM", "PG", "PM", "SG", "SM", "KG", "KM", "RG", "RM"},
		},
	}

	path := bfs.Find(initialState, endState.ID())
	fmt.Printf("Steps: %d\n", len(path)-1)
}

func part2() {
	initialState := &State{
		elevator: 0,
		levels: [][]string{
			{"TG", "TM", "PG", ".", "SG", ".", ".", ".", ".", ".", "EG", "EM", "DG", "DM"},
			{".", ".", ".", "PM", ".", "SM", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", "KG", "KM", "RG", "RM", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		},
	}

	endState := &State{
		elevator: 3,
		levels: [][]string{
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
			{"TG", "TM", "PG", "PM", "SG", "SM", "KG", "KM", "RG", "RM", "EG", "EM", "DG", "DM"},
		},
	}

	path := bfs.Find(initialState, endState.ID())
	fmt.Printf("Steps: %d\n", len(path)-1)
}

func main() {
	part1()
	part2()
}
