package main

import (
	"fmt"
	"github.com/abates/AdventOfCode/2016/util"
	"strings"
)

func newWord(length int) []*util.Letters {
	word := make([]*util.Letters, length)
	for i := 0; i < length; i++ {
		word[i] = util.NewLetters()
	}

	return word
}

func main() {
	var word []*util.Letters
	for _, line := range util.ReadInput() {
		if word == nil {
			word = newWord(len(line))
		}

		for i, l := range strings.Split(line, "") {
			word[i].Add(l)
		}
	}

	for _, letters := range word {
		l := letters.Sorted()
		fmt.Printf("%s", l[len(l)-1])
	}
	fmt.Printf("\n")
}
