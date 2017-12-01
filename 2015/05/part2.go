package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func pairs(word string) bool {
	pairs := make(map[string]int)

	for i := 0; i < len(word)-1; i++ {
		pair := word[i : i+2]
		pairs[pair]++
		if pairs[pair] >= 2 {
			fmt.Printf("Word %s pair %s repeats at position %d\n", word, pair, i)
			return true
		}

		if i < len(word)-2 {
			if pair == word[i+1:i+3] {
				i++
			}
		}
	}
	return false
}

func repeat(word string) bool {
	var last1 rune
	var last2 rune
	for i, r := range word {
		if i > 1 && last2 == r {
			return true
		}
		last2 = last1
		last1 = r
	}
	return false
}

func isNice(word string) bool {
	for _, t := range []func(string) bool{pairs, repeat} {
		if !t(word) {
			return false
		}
	}
	return true
}

func main() {
	//f, _ := os.Open("input.txt")
	//b, _ := ioutil.ReadAll(f)
	b, _ := ioutil.ReadAll(os.Stdin)
	nice := 0
	for _, line := range strings.Split(string(b), "\n") {
		if line == "" {
			continue
		}
		if isNice(line) {
			nice++
		}
	}
	fmt.Println(nice)
}
