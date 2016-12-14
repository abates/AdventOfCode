package main

import ()

/*
 * The first floor contains a thulium generator, a thulium-compatible microchip, a plutonium generator, and a strontium generator.
 * The second floor contains a plutonium-compatible microchip and a strontium-compatible microchip.
 * The third floor contains a promethium generator, a promethium-compatible microchip, a ruthenium generator, and a ruthenium-compatible microchip.
 * The fourth floor contains nothing relevant.
 */
func main() {
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

	initialState.Find(endState, 0)
}
