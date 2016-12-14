package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
	"strconv"
)

const blankItem = "."

func CanLive(items []string) bool {
	if len(items) == 0 || len(items) == 1 {
		return true
	}

	unbalancedGenerators := false
	unbalancedMicrochips := false

	for i := 0; i < len(items); i++ {
		if items[i] == blankItem {
			continue
		}

		if items[i][1] == 'G' {
			if i+1 < len(items) && items[i][0] == items[i+1][0] {
				i++
			} else {
				unbalancedGenerators = true
			}
		} else {
			unbalancedMicrochips = true
		}
	}

	return !unbalancedGenerators || !unbalancedMicrochips
}

func pairs(level []string) [][]int {
	set := make([][]int, 0)

	for i, item1 := range level {
		if item1 == blankItem {
			continue
		}

		for j, item2 := range level[i+1:] {
			if item1 == item2 || item2 == blankItem {
				continue
			}
			set = append(set, []int{i, i + j + 1})
		}
		set = append(set, []int{i})
	}
	return set
}

type Hash uint64
type Level int

type State struct {
	elevator      Level
	levels        [][]string
	visitedStates map[Hash]bool
	nextStates    map[Hash]*State
	hash          Hash
}

func ReadStateFromHash(hash string) *State {
	newState := &State{}

	elevator, _ := strconv.Atoi(hash[0:1])
	newState.elevator = Level(elevator)
	hash = hash[1:]
	items := make([]string, 0)
	for i := 0; i < len(hash); i++ {
		if hash[i] == '.' {
			items = append(items, hash[i:i+1])
		} else {
			items = append(items, hash[i:i+2])
			i++
		}
	}

	length := len(items) / 4
	for i := 0; i < 4; i++ {
		start := length * i
		newState.levels = append(newState.levels, items[start:start+length])
	}
	return newState
}

func NewState(oldState *State) *State {
	newState := &State{
		elevator:      oldState.elevator,
		levels:        make([][]string, len(oldState.levels)),
		visitedStates: oldState.visitedStates,
	}

	for i, level := range oldState.levels {
		newState.levels[i] = make([]string, len(level))
		copy(newState.levels[i], level)
	}
	return newState
}

func (s *State) Move(fromItem int, toFloor Level) {
	s.levels[toFloor][fromItem] = s.levels[s.elevator][fromItem]
	s.levels[s.elevator][fromItem] = blankItem
}

func (s *State) HashString() string {
	writer := util.StringWriter{}

	writer.Writef("%d", s.elevator)
	for _, level := range s.levels {
		for _, item := range level {
			writer.Write(item)
		}
	}
	return writer.String()
}

func (s *State) Hash() Hash {
	hash := Hash(s.elevator)
	for _, level := range s.levels {
		for _, item := range level {
			hash = hash << 1
			if item != blankItem {
				hash |= 0x01
			}
		}
	}
	return hash
}

func (s *State) Equal(other *State) bool {
	return s.Hash() == other.Hash()
}

func (s *State) String() string {
	writer := &util.StringWriter{}
	for i := len(s.levels) - 1; i >= 0; i-- {
		items := s.levels[i]
		if Level(i) == s.elevator {
			writer.Write("E")
		} else {
			writer.Write(blankItem)
		}
		for _, item := range items {
			writer.Writef("%2s ", item)
		}
		writer.Write("\n")
	}
	return writer.String()
}

func (s *State) createNextState(level Level, itemsToMove []int) (newState *State) {
	newState = NewState(s)
	for _, itemIndex := range itemsToMove {
		newState.Move(itemIndex, level)
	}
	newState.elevator = level

	return newState
}

func (s *State) NextStates() map[Hash]*State {
	if s.nextStates == nil {
		s.nextStates = make(map[Hash]*State, 0)
		for _, pair := range pairs(s.levels[s.elevator]) {
			if len(pair) == 2 {
				item1 := s.levels[s.elevator][pair[0]]
				item2 := s.levels[s.elevator][pair[1]]
				if !CanLive([]string{item1, item2}) {
					continue
				}
			}

			for _, direction := range []int{1, -1} {
				delta := Level(direction)
				if s.elevator+delta < Level(len(s.levels)) && s.elevator+delta >= 0 {
					nextState := s.createNextState(s.elevator+delta, pair)
					nextHash := nextState.Hash()
					if _, found := s.visitedStates[nextHash]; !found {
						if CanLive(nextState.levels[nextState.elevator]) {
							s.visitedStates[nextHash] = true
							s.nextStates[nextHash] = nextState
						} else {
							break
						}
					}
				}
			}
		}
	}
	return s.nextStates
}

func (s *State) Find(endState *State, depth int) {
	s.visitedStates = make(map[Hash]bool)
	for ; ; depth++ {
		fmt.Printf("Depth %d\n", depth)
		if path, moreBranches := s.find(endState, depth, nil); path != nil {
			/*for _, p := range path {
				fmt.Printf("%s--------------------------------------\n", p)
			}
			fmt.Printf("==========================================\n")*/
			fmt.Printf("%d\n", depth)
			break
		} else if moreBranches == false {
			break
		}
	}
}

func (s *State) find(endState *State, depth int, path []*State) ([]*State, bool) {
	nextStates := s.NextStates()
	if depth == 0 {
		if s.Equal(endState) {
			return path, false
		}
		return nil, len(nextStates) > 0
	}

	for _, nextState := range nextStates {
		path = append(path, nextState)
		p, b := nextState.find(endState, depth-1, path)
		if p != nil {
			return p, false
		}

		if b == false {
			delete(s.nextStates, nextState.Hash())
		}
		path = path[:len(path)-1]
	}
	return nil, len(nextStates) > 0
}
