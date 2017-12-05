package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Word string

func anagram(input Word) []Word {
	words := make([]Word, 0)
	if len(input) == 1 {
		return []Word{input}
	}

	if len(input) > 0 {
		l := input[0:1]
		input = input[1:]
		for _, word := range anagram(input) {
			for i := 0; i <= len(word); i++ {
				words = append(words, Word(fmt.Sprintf("%s%s%s", word[0:i], l, word[i:])))
			}
		}
	}
	return words
}

func (w Word) Anagrams() []Word {
	return anagram(w)
}

type PassPhrase []Word

func NewPassPhrase(phrase string) PassPhrase {
	passPhrase := make([]Word, 0)
	for _, word := range strings.Fields(phrase) {
		passPhrase = append(passPhrase, Word(word))
	}
	return passPhrase
}

func (p PassPhrase) Valid() bool {
	words := make(map[Word]bool)
	for _, w1 := range p {
		if _, found := words[w1]; found {
			return false
		}

		words[w1] = true
		for _, w2 := range w1.Anagrams() {
			if w1 == w2 {
				continue
			}
			if _, found := words[w2]; found {
				return false
			}
		}
	}
	return true
}

func main() {
	valid := 0
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		phrase := NewPassPhrase(line)
		if phrase.Valid() {
			valid++
		}
	}
	fmt.Printf("Valid Phrases: %d\n", valid)
}
