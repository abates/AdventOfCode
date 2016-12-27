/**
* The method of generating states was adapted from:
 *       Author: Andrew Foote
 *         Date: 2016-12-11
 * Availability: https://andars.github.io/aoc_day11.html
*/
package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/bfs"
	"github.com/abates/AdventOfCode/2016/util"
	"sort"
	"strconv"
	"strings"
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
	elevator Level
	levels   [][]string
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
		elevator: oldState.elevator,
		levels:   make([][]string, len(oldState.levels)),
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

func (s *State) ID() string {
	pairs := make([]string, 0)
	items := s.levels[0]
	pair := ""
	for i := 0; i < len(items); i++ {
		if i%2 == 0 {
			pair = "("
		}
		for j := 0; j < len(s.levels); j++ {
			if s.levels[j][i] != blankItem {
				pair = fmt.Sprintf("%s%d", pair, j)
				break
			}
		}
		if i%2 == 0 {
			pair = fmt.Sprintf("%s,", pair)
		} else {
			pair = fmt.Sprintf("%s)", pair)
			pairs = append(pairs, pair)
		}
	}
	sort.Strings(pairs)
	return fmt.Sprintf("%d,%s", s.elevator, strings.Join(pairs, ","))
}

func (s *State) Equal(node bfs.Node) bool {
	if other, ok := node.(*State); ok {
		return s.ID() == other.ID()
	}
	return false
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

func (s *State) Neighbors() []bfs.Node {
	nodes := make([]bfs.Node, 0)
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
				if CanLive(nextState.levels[nextState.elevator]) {
					nodes = append(nodes, nextState)
				} else {
					break
				}
			}
		}
	}
	return nodes
}
