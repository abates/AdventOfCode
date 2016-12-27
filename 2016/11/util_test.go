package main

import (
	"reflect"
	"testing"
)

func TestCanLive(t *testing.T) {
	tests := []struct {
		result bool
		items  []string
	}{
		{true, []string{"HG"}},
		{true, []string{"HG", "."}},
		{true, []string{"HG", "HM"}},
		{true, []string{"HG", "HM", "."}},
		{true, []string{"HG", "HM", "LG"}},
		{true, []string{"HG", "HM", "LM"}},
		{false, []string{"HM", "LG"}},
		{false, []string{"HM", "LG", "."}},
		{true, []string{"HG", "LG"}},
		{true, []string{"HG", "LG", "."}},
		{true, []string{"HG", "LG", "RG"}},
		{true, []string{"HG", "LG", "RG", "."}},
		{false, []string{"HG", "LG", "LM", "RM"}},
		{false, []string{"HG", "LG", "LM", "RM", "."}},
		{true, []string{"HG", "HM", "LG", "LM", "RG"}},
		{true, []string{"HG", "HM", "LG", "LM", "RG", "."}},
	}

	for i, test := range tests {
		result := CanLive(test.items)
		if result != test.result {
			t.Errorf("Test %d failed.  Expected %v got %v", i, test.result, result)
		}
	}
}

func TestPairs(t *testing.T) {
	tests := []struct {
		input  []string
		result [][]int
	}{
		{[]string{".", "1", ".", "."}, [][]int{[]int{1}}},
		{[]string{".", "1", "2", "."}, [][]int{[]int{1, 2}, []int{1}, []int{2}}},
		{[]string{".", "1", "2", "3"}, [][]int{[]int{1, 2}, []int{1, 3}, []int{1}, []int{2, 3}, []int{2}, []int{3}}},
	}

	for i, test := range tests {
		result := pairs(test.input)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("Test %d failed.  Expected %v got %v", i, test.result, result)
		}
	}
}

func TestHashString(t *testing.T) {
	state := &State{
		elevator: 0,
		levels: [][]string{
			{".", "HM", ".", "LM"},
			{"HG", ".", ".", "."},
			{".", ".", "LG", "."},
			{".", ".", ".", "."},
		},
	}

	expected := "0.HM.LMHG.....LG....."
	if state.HashString() != expected {
		t.Errorf("Expected %s got %s", expected, state.HashString())
	}

	newState := ReadStateFromHash(expected)
	if newState.HashString() != expected {
		t.Errorf("Expected %s got %s", expected, newState.HashString())
	}
}

func TestConstructor(t *testing.T) {
	state := ReadStateFromHash("0.HM.LMHG.....LG.....")
	newState := NewState(state)
	if !state.Equal(newState) {
		t.Errorf("Constructor failed.  Expected %s got %s", state, newState)
	}
}

func TestMove(t *testing.T) {
	startState := ReadStateFromHash("0.HM.LMHG.....LG.....")
	expectedState := ReadStateFromHash("0...LMHG.....LG..HM..")

	startState.Move(1, 3)
	if !startState.Equal(expectedState) {
		t.Errorf("Constructor failed.  Expected %s got %s", startState, expectedState)
	}
}

func TestNextStates(t *testing.T) {
	tests := []struct {
		initialState *State
		validStates  []string
	}{
		{
			initialState: ReadStateFromHash("0.HM.LMHG.....LG....."),
			validStates: []string{
				"1....HGHM.LM..LG.....",
				//"2....HG....HMLGLM....",
				//"3....HG.....LG..HM.LM",
				"1...LMHGHM....LG.....",
			},
		}, {
			initialState: ReadStateFromHash("1...LMHGHM....LG....."),
			validStates: []string{
				"2...LM....HGHMLG.....",
				//"3...LM......LG.HGHM..",
				"0HGHM.LM......LG.....",
				"2...LM.HM..HG.LG.....",
				//"3...LM.HM....LG.HG...",
				//"0.HM.LMHG.....LG.....",
			},
		},
	}
	//F3 .  .  .  .  .
	//F2 .  .  .  LG .
	//F1 E  HG HM .  .
	//F0 .  .  .  .  LM

	for i, test := range tests {
		validStates := make([]*State, 0)
		for _, hashString := range test.validStates {
			state := ReadStateFromHash(hashString)
			validStates = append(validStates, state)
		}

		nextStates := test.initialState.Neighbors()
		if len(nextStates) != len(test.validStates) {
			t.Errorf("Test %d Expected %d next states.  Got %d", i, len(test.validStates), len(nextStates))
		}

		if len(nextStates) != len(validStates) {
			t.Errorf("Test %d Expected next states %v got %v", i, validStates, nextStates)
		} else {
			for j, state := range validStates {
				if !state.Equal(nextStates[j]) {
					t.Errorf("Test %d Expected state\n%s\nbut got\n%s", i, state, nextStates[j])
				}
			}
		}
	}
}
