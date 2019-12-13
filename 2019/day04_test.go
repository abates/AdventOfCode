package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestD4Parse(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		wantMin int
		wantMax int
	}{
		{"test 1", []string{"1234-4567"}, 1234, 4567},
	}

	for _, test := range tests {
		t.Run("Parsing "+test.name, func(t *testing.T) {
			d4 := &D4{}
			err := d4.parseFile(test.input)
			if err == nil {
			} else {
				t.Errorf("Unexpected Error: %v", err)
			}
		})
	}
}

func TestD4GetParts(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{111111, []int{1, 1, 1, 1, 1, 1}},
		{223450, []int{2, 2, 3, 4, 5, 0}},
		{123789, []int{1, 2, 3, 7, 8, 9}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			got := splitInt(6, test.input)
			if !reflect.DeepEqual(test.want, got) {
				t.Errorf("Wanted %v got %v", test.want, got)
			}
		})
	}
}

func TestD4CheckIncrementing(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{111111, true},
		{223450, false},
		{123789, true},
		{177777, true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			got := checkIncrementing(splitInt(6, test.input))
			if test.want != got {
				t.Errorf("Wanted %v got %v", test.want, got)
			}
		})
	}
}

func TestD4CheckRepeating(t *testing.T) {
	tests := []struct {
		name       string
		input      int
		doubleOnly bool
		want       bool
	}{
		{"111111", 111111, false, true},
		{"223450", 223450, false, true},
		{"123789", 123789, false, false},
		{"177777", 177777, false, true},
		{"111111 (unique doubles)", 111111, true, false},
		{"223450 (unique doubles)", 223450, true, true},
		{"123789 (unique doubles)", 123789, true, false},
		{"177777 (unique doubles)", 177777, true, false},
		{"112233 (unique doubles)", 112233, true, true},
		{"123444 (unique doubles)", 123444, true, false},
		{"111122 (unique doubles)", 111122, true, true},
		{"177999 (unique doubles)", 177999, true, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := checkRepeating(splitInt(6, test.input), test.doubleOnly)
			if test.want != got {
				t.Errorf("Wanted %v got %v", test.want, got)
			}
		})
	}
}

func TestD4Parts(t *testing.T) {
	tests := []challengeTest{}

	for _, test := range tests {
		d4 := &D4{}
		challenge := &challenge{"Test Day 04", "", nil, d4.parseFile, d4.part1, d4.part2}
		t.Run("Parsing "+test.name, testChallenge(challenge, test))
	}
}
