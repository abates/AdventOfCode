package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type PassPhrase string

func (p PassPhrase) Valid() bool {
	words := make(map[string]bool)
	for _, word := range strings.Fields(string(p)) {
		if _, found := words[word]; found {
			return false
		}
		words[word] = true
	}
	return true
}

func main() {
	valid := 0
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		phrase := PassPhrase(line)
		if phrase.Valid() {
			fmt.Printf("Valid: %q\n", phrase)
			valid++
		}
	}
	fmt.Printf("Valid Phrases: %d\n", valid)
}
