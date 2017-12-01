package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func vowels(word string) bool {
	count := 0
	for _, r := range word {
		if strings.ContainsRune("aeiou", r) {
			count++
		}
	}
	return count >= 3
}

func repeat(word string) bool {
	var last rune
	for i, r := range word {
		if i > 0 && last == r {
			return true
		}
		last = r
	}
	return false
}

func badCombos(word string) bool {
	for _, str := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(word, str) {
			return false
		}
	}
	return true
}

func isNice(word string) bool {
	for _, t := range []func(string) bool{vowels, repeat, badCombos} {
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
