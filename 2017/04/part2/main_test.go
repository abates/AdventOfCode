package main

import (
	"reflect"
	"testing"
)

func TestAnagram(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "bac", "bca", "acb", "cab", "cba"}},
	}

	for i, test := range tests {
		anagrams := anagram(Word(test.input))
		words := make([]string, len(anagrams))
		for i, word := range anagrams {
			words[i] = string(word)
		}
		if !reflect.DeepEqual(test.expected, words) {
			t.Errorf("tests[%d] expected %-v got %-v", i, test.expected, words)
		}
	}
}

func TestValid(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
		{"kvvfl kvvfl olud wjqsqa olud frc", false},
	}

	for i, test := range tests {
		phrase := NewPassPhrase(test.input)
		if phrase.Valid() != test.valid {
			t.Errorf("tests[%d] expected %v got %v", i, test.valid, phrase.Valid())
		}
	}
}
