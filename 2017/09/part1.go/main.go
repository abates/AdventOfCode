package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func parseGarbage(reader *strings.Reader) (garbage int) {
	for 0 < reader.Len() {
		r, _, _ := reader.ReadRune()
		switch r {
		case '!':
			reader.ReadRune()
		case '>':
			return
		default:
			garbage++
		}
	}
	return
}

type Group struct {
	Children []*Group
}

func (g *Group) Score(parentScore int) int {
	parentScore += 1
	score := 0
	for _, child := range g.Children {
		childScore := child.Score(parentScore)
		score += childScore
	}
	return parentScore + score
}

func parseGroup(reader *strings.Reader) (group *Group, garbage int) {
	for 0 < reader.Len() {
		r, _, _ := reader.ReadRune()
		switch r {
		case '{':
			if group == nil {
				group = &Group{}
			} else {
				reader.UnreadRune()
				newGroup, newGarbage := parseGroup(reader)
				garbage += newGarbage
				group.Children = append(group.Children, newGroup)
			}
		case '<':
			garbage += parseGarbage(reader)
		case '}':
			return group, garbage
		}
	}
	return group, garbage
}

func main() {
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	str := string(b)

	group, garbage := parseGroup(strings.NewReader(str))
	fmt.Printf("Count: %d\n", group.Score(0))
	fmt.Printf("Garbage: %d\n", garbage)
}
